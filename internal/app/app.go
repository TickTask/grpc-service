package app

import (
	"log/slog"
	grpcapp "server/internal/app/grpc"
	"server/internal/services/tasks"
	"server/internal/services/user"
	"server/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}
	userService := user.New(log, storage, storage, storage, accessTokenTTL, refreshTokenTTL)
	tasksService := tasks.New(log)
	grpcApp := grpcapp.New(grpcPort, log, userService, tasksService)
	return &App{
		grpcApp,
	}
}
