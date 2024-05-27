package configs

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func initLogger(logLevelStr string) {
	logLevel := getLogLevel(logLevelStr)

	w := os.Stderr
	slog.SetDefault(
		slog.New(tint.NewHandler(w, &tint.Options{
			Level:      logLevel,
			TimeFormat: time.RFC3339,
		})),
	)

}

func getLogLevel(s string) slog.Level {
	switch s {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}

}
