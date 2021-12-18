// Package logger expose all log levels for the api
package logger

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// SetLoggerLogLevel permit to set default log level
func SetLoggerLogLevel() {
	switch strings.TrimSpace(os.Getenv("APP_LOG_LEVEL")) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		gin.SetMode("debug")
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
