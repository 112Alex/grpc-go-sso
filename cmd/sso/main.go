package main

import (
	"log/slog"
	"os"

	"github.com/112Alex/grpc-go-sso/SSO/internal/app"
	"github.com/112Alex/grpc-go-sso/SSO/internal/config"
	"github.com/112Alex/grpc-go-sso/SSO/lib/logger/handlers/slogpretty"
)

const (
	envLoocal = "local"
	envDev    = "dev"
	envProd   = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg)) //NOTE: must edit on prod

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCSrv.MustRun()

	// TODO: initialize app
	// NOTE: The application itself, not the entry point

	//TODO: launch the gRPC server of the application.
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLoocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
