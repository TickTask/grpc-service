package tasks

import (
	"context"
	"fmt"
	"log/slog"
	"server/internal/domain/model"
	"time"
)

type Task struct {
	log          *slog.Logger
	saverTask    SaverTask
	removerTask  RemoverTask
	providerTask ProviderTask
}

func New(
	log *slog.Logger,
	saverTask SaverTask,
	removerTask RemoverTask,
	providerTask ProviderTask,
) *Task {
	return &Task{
		log:          log,
		saverTask:    saverTask,
		removerTask:  removerTask,
		providerTask: providerTask,
	}
}

type SaverTask interface {
	SaveTask(ctx context.Context, task model.RequestTask) (int64, error)
}

type RemoverTask interface {
	Remove(ctx context.Context, taskID int64) error
}

type ProviderTask interface {
	GetTaskByID(ctx context.Context, taskID int64) (model.Task, error)
}

func (t *Task) CreateTask(ctx context.Context, title string, body string) (int64, error) {
	const op = "task.create"

	userID, ok := ctx.Value("user_id").(int64)

	if !ok {
		return 0, fmt.Errorf("Not found user_id in context")
	}

	var task model.RequestTask

	task.Title = title

	task.Body = body

	task.CreatedAt = time.Now().UTC()

	task.UserID = userID

	task.StatusID = 1

	id, err := t.saverTask.SaveTask(ctx, task)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (t *Task) FetchTask(ctx context.Context, taskID int64) (model.Task, error) {
	return model.Task{}, nil
}

func (t *Task) RemoveTask(ctx context.Context, taskID int64) error {
	const op = "task.remove"

	err := t.removerTask.Remove(ctx, taskID)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
