package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/112Alex/grpc-go-sso/SSO/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	StoragePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: initialize storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
