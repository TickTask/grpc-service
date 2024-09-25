package tasks

import (
	"context"
	"google.golang.org/grpc"
	taskrpc "server/pkg/task"
)

type serverApi struct {
	taskrpc.UnimplementedTaskServer
	tasks Tasks
}

func Register(gRPC *grpc.Server, tasks Tasks) {
	taskrpc.RegisterTaskServer(gRPC, &serverApi{tasks: tasks})
}

type Tasks interface {
}

func (s *serverApi) CreateTask(ctx context.Context, in *taskrpc.CreateTaskRequest) (*taskrpc.CreateTaskResponse, error) {
	return &taskrpc.CreateTaskResponse{}, nil
}
