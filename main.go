package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Lord-Y/versions/commons"
	customLogger "github.com/Lord-Y/versions/logger"
	"github.com/Lord-Y/versions/mysql"
	"github.com/Lord-Y/versions/postgres"
	"github.com/Lord-Y/versions/routers"
	"github.com/rs/zerolog/log"
)

//go:embed sql/postgres
var sqlpostgres embed.FS

//go:embed sql/mysql
var sqlmysql embed.FS

func init() {
	customLogger.SetLoggerLogLevel()
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
		mysql.InitDB(sqlmysql)
	case "postgres":
		postgres.InitDB(sqlpostgres)
	default:
		msg := "SQL_DRIVER environment variable can only be mysql or postgres"
		log.Fatal().Err(fmt.Errorf(msg)).Msg(msg)
		return
	}
}

func main() {
	var appPort string
	router := routers.SetupRouter()
	port := strings.TrimSpace(os.Getenv("APP_PORT"))

	switch strings.HasPrefix(port, ":") {
	case true:
		appPort = port
	case false:
		if port == "" {
			appPort = ":8080"
		} else {
			appPort = fmt.Sprintf(":%s", port)
		}
	}

	srv := &http.Server{
		Addr:    appPort,
		Handler: router,
	}
	log.Info().Msgf("Starting server on port %s", appPort)

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
	log.Info().Msg("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server shutted down abruptly")
	}
	log.Info().Msg("Server exited successfully")
}
