package main

import (
	"log/slog"
	"os"

	"github.com/112Alex/grpc-go-sso/SSO/internal/config"
)

const (
	envLoocal = "local"
	envDev    = "dev"
	envProd   = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
	)

	log.Warn("warning message")

	// TODO: initialize logger (slog)

	// TODO: initialize app
	// NOTE: Само приложение, а не точка входа
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLoocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
