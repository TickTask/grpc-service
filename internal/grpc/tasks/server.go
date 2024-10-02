package tasks

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"server/internal/domain/model"
	"server/internal/lib/mapper"
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
	CreateTask(ctx context.Context, title string, body string) (int64, error)
	RemoveTask(ctx context.Context, taskID int64) error
	FetchTask(ctx context.Context, taskID int64) (model.Task, error)
	FetchTasks(ctx context.Context) ([]model.Task, error)
}

func (s *serverApi) CreateTask(ctx context.Context, request *taskrpc.CreateTaskRequest) (*taskrpc.CreateTaskResponse, error) {
	err := validateCreateTaskRequest(request)

	if err != nil {
		return nil, err
	}

	id, err := s.tasks.CreateTask(ctx, request.GetTitle(), request.GetBody())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &taskrpc.CreateTaskResponse{
		TaskId: id,
	}, nil
}

func (s *serverApi) GetTask(ctx context.Context, request *taskrpc.GetTaskRequest) (*taskrpc.GetTaskResponse, error) {
	task, err := s.tasks.FetchTask(ctx, request.GetTaskId())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return mapper.ToTaskResponse(task), nil
}

func (s *serverApi) DeleteTask(ctx context.Context, request *taskrpc.DeleteTaskRequest) (*emptypb.Empty, error) {
	err := validateDeleteTaskRequest(request)

	if err != nil {
		return nil, err
	}

	err = s.tasks.RemoveTask(ctx, request.GetTaskId())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *serverApi) GetTasks(ctx context.Context, _ *emptypb.Empty) (*taskrpc.GetTasksResponse, error) {
	tasks, err := s.tasks.FetchTasks(ctx)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response := mapper.ToTasksResponse(tasks)

	return &taskrpc.GetTasksResponse{Tasks: response}, nil
}

func validateCreateTaskRequest(request *taskrpc.CreateTaskRequest) error {

	if request.GetTitle() == "" {
		return status.Error(codes.InvalidArgument, "title is required")
	}

	return nil
}

func validateDeleteTaskRequest(request *taskrpc.DeleteTaskRequest) error {
	if request.GetTaskId() == 0 {
		return status.Error(codes.InvalidArgument, "task id is required")
	}
	return nil
}
