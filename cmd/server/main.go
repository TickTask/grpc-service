package main

import (
	"log/slog"
	"os"
	"os/signal"
	"server/internal/app"
	"server/internal/config"
	"server/internal/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCServer.Stop()

	log.Info("stopping application")

}
