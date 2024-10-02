package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"server/internal/domain/model"
	"server/pkg/status"
	"server/pkg/task"
	"server/pkg/user"
)

func ToTaskResponse(model model.Task) *task.GetTaskResponse {
	u := &user.UserData{
		UserId:   model.User.ID,
		Username: model.User.Name,
		Login:    model.User.Login,
	}

	s := &status.StatusData{
		Id:     model.Status.ID,
		Status: model.Status.Status,
	}

	t := &task.GetTaskResponse{
		TaskId:   model.ID,
		Title:    model.Title,
		Body:     model.Body,
		CreateAt: timestamppb.New(model.CreatedAt),
		User:     u,
		Status:   s,
	}
	return t
}

func ToTasksResponse(tasks []model.Task) []*task.GetTaskResponse {
	var taskResponse []*task.GetTaskResponse
	for _, t := range tasks {
		taskResponse = append(taskResponse, ToTaskResponse(t))
	}

	return taskResponse
}
