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
	customLogger "github.com/Lord-Y/versions-api/logger"
	"github.com/Lord-Y/versions-api/mysql"
	"github.com/Lord-Y/versions-api/postgres"
	"github.com/Lord-Y/versions-api/routers"
	"github.com/rs/zerolog/log"
)

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
	router := routers.SetupRouter()
	appPort := strings.TrimSpace(os.Getenv("APP_PORT"))
	if appPort != "" {
		srv = &http.Server{
			Addr:    fmt.Sprintf(":%s", appPort),
			Handler: router,
		}
		log.Info().Msgf("Starting server on port %s", appPort)
	} else {
		appPort = ":8080"
		srv = &http.Server{
			Addr:    appPort,
			Handler: router,
		}
		log.Info().Msgf("Starting server on port %s", appPort)
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
	log.Info().Msg("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server shutted down abruptly")
	}
	log.Info().Msg("Server exited successfully")
}
