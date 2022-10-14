package routers

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/Lord-Y/versions/health"
	customLogger "github.com/Lord-Y/versions/logger"
	"github.com/Lord-Y/versions/metrics"
	"github.com/Lord-Y/versions/versionning"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	customLogger.SetLoggerLogLevel()
}

//go:embed ui/dist/assets
var assets embed.FS

//go:embed ui
var ui_ embed.FS

// SetupRouter gin
func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	router := gin.New()
	router.Use(gin.Recovery())
	router.RedirectTrailingSlash = true

	router.Use(
		logger.SetLogger(
			logger.WithUTC(true),
			logger.WithLogger(
				func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
					var d time.Duration
					return zerolog.New(os.Stdout).
						With().
						Timestamp().
						Int("status", c.Writer.Status()).
						Str("method", c.Request.Method).
						Str("path", c.Request.URL.Path).
						Str("ip", c.ClientIP()).
						Dur("latency", d).
						Str("user_agent", c.Request.UserAgent()).
						Logger()
				},
			),
		),
	)

	// disable during unit testing
	if os.Getenv("APP_PROMETHEUS") != "" {
		p := ginprometheus.NewPrometheus("http")
		p.SetListenAddressWithRouter(":9101", router)
		p.Use(router)
		prometheus.MustRegister(metrics.LastDeployments())
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", health.Health)
		v1.HEAD("/health", health.Health)
		v1.GET("/healthz", health.Healthz)
		v1.HEAD("/healthz", health.Healthz)

		v1.POST("/create", versionning.Create)
		v1.POST("/update/status", versionning.UpdateStatus)
		v1.GET("/read/environment", versionning.ReadEnvironment)
		v1.GET("/read/platform", versionning.ReadPlatform)
		v1.GET("/read/home", versionning.ReadHome)
		v1.GET("/read/distinct/workloads", versionning.ReadDistinctWorkloads)
		v1.GET("/read/raw", versionning.Raw)
		v1.GET("/read/raw/id", versionning.RawById)
		v1.GET("/read/environment/latest", versionning.ReadEnvironmentLatest)
		v1.GET("/read/environment/latest/whatever", versionning.ReadEnvironmentLatest)
		v1.GET("/stats/latest", versionning.GetLastXDaysDeployments)
	}

	router.StaticFS("/ui/assets", EmbedFolder(assets, "ui/dist/assets"))

	// done like that to avoid trailing slash
	router.GET("/ui/logo.png", func(c *gin.Context) {
		f, _ := ui_.ReadFile("ui/dist/logo.png")
		c.Data(
			http.StatusOK,
			"image/png",
			f,
		)
	})

	// done like that to avoid trailing slash
	router.GET("/ui/favicon.ico", func(c *gin.Context) {
		f, _ := ui_.ReadFile("ui/dist/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			f,
		)
	})

	ui := router.Group("/ui", func(c *gin.Context) {
		f, _ := ui_.ReadFile("ui/dist/index.html")
		c.Data(
			http.StatusOK,
			"text/html",
			f,
		)
	})
	if ui.BasePath() == "/ui" {
		router.NoRoute(func(c *gin.Context) {
			f, _ := ui_.ReadFile("ui/dist/index.html")
			c.Data(
				http.StatusOK,
				"text/html",
				f,
			)
		})
	}

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/ui/")
	})

	return router
}

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	s, _ := f.Stat()
	if s.IsDir() && !e.indexes {
		return false
	}

	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
