package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"server/internal/grpc/tasks"
	"server/internal/grpc/user"
	"server/internal/lib/interceptors"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(port int, log *slog.Logger, userService user.User, tasksService tasks.Tasks) *App {
	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(interceptors.IsAuth))

	user.Register(gRPCServer, userService)

	tasks.Register(gRPCServer, tasksService)
	
	return &App{
		port:       port,
		gRPCServer: gRPCServer,
		log:        log,
	}
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}

func (a *App) run() error {
	const op = "grpcapp.run"

	log := a.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server on port", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.stop"
	a.log.With(slog.String("op", op)).Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
