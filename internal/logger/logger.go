package logger

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

const (
	LOCAL = "local"
	DEV   = "dev"
	PROD  = "prod"
)

func New(env string) (*slog.Logger, error) {
	var log *slog.Logger
	switch env {
	case LOCAL:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case DEV:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case PROD:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		return nil, fmt.Errorf("invalid logger env")
	}

	return log, nil
}
