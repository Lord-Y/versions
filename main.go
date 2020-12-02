package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Lord-Y/versions-api/commons"
	"github.com/Lord-Y/versions-api/health"
	"github.com/Lord-Y/versions-api/mysql"
	"github.com/Lord-Y/versions-api/postgres"
	"github.com/Lord-Y/versions-api/versionning"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetupRouter gin
func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	switch strings.TrimSpace(os.Getenv("APP_LOG_LEVEL")) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		gin.SetMode("debug")
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	subLog := zerolog.New(os.Stdout).With().Timestamp().Logger()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger.SetLogger(logger.Config{
		Logger: &subLog,
		UTC:    true,
	}))

	// disable during unit testing
	if os.Getenv("APP_PROMETHEUS") != "" {
		p := ginprometheus.NewPrometheus("http")
		p.SetListenAddress(":9101")
		p.Use(router)
	}

	v1 := router.Group("/api/v1/versions")
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
	}
	return router
}

func init() {
	if commons.RedisEnabled() {
		if commons.GetRedisURI() == "" {
			msg := "REDIS_URI environment variable must be set"
			log.Fatal().Err(fmt.Errorf(msg)).Msg(msg)
			return
		}
	}

	if os.Getenv("SLEEP") != "" {
		sleep, _ := strconv.Atoi(os.Getenv("SLEEP"))
		log.Info().Msgf("Sleeping %d", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	if commons.SqlDriver == "" {
		msg := "SQL_DRIVER environment variable can only be mysql or postgres"
		log.Fatal().Err(fmt.Errorf(msg)).Msg(msg)
		return
	}
	if strings.TrimSpace(os.Getenv("DB_URI")) == "" {
		msg := "DB_URI environment variable must be set"
		log.Fatal().Err(fmt.Errorf(msg)).Msg(msg)
		return
	}

	switch commons.SqlDriver {
	case "mysql":
		mysql.InitDB()
	case "postgres":
		postgres.InitDB()
	default:
		msg := "SQL_DRIVER environment variable can only be mysql or postgres"
		log.Fatal().Err(fmt.Errorf(msg)).Msg(msg)
		return
	}
}

func main() {
	var srv *http.Server
	router := SetupRouter()
	appPort := strings.TrimSpace(os.Getenv("APP_PORT"))
	if appPort != "" {
		srv = &http.Server{
			Addr:    fmt.Sprintf(":%s", appPort),
			Handler: router,
		}
	} else {
		srv = &http.Server{
			Addr:    ":8080",
			Handler: router,
		}
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Startup failed")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server Shutdown")
	}
	log.Info().Msg("Server exiting")
}
