package tasks

import "log/slog"

type Task struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Task {
	return &Task{log: log}
}
