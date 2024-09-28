package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/domain/model"
)

type TaskStorage struct {
	db *sql.DB
}

func NewTaskStorage(storagePath string) (*TaskStorage, error) {
	const op = "storage.sqlite.new"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &TaskStorage{db: db}, nil
}

func (t *TaskStorage) Stop() error {
	return t.db.Close()
}

func (t *TaskStorage) SaveTask(ctx context.Context, task model.RequestTask) (int64, error) {
	const op = "storage.sqlite.save_task"

	req, err := t.db.Prepare("INSERT INTO Tasks(title, body, created_at, task_user_id, task_status_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := req.ExecContext(ctx, task.Title, task.Body, task.CreatedAt, task.UserID, task.StatusID)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (t *TaskStorage) GetTaskByID(ctx context.Context, taskID int64) (model.Task, error) {

	return model.Task{}, nil
}

func (t *TaskStorage) Remove(ctx context.Context, taskID int64) error {
	const op = "storage.sqlite.remove_task"
	req, err := t.db.Prepare("DELETE FROM Tasks WHERE id = ?")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := req.ExecContext(ctx, taskID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: no rows affected", op)
	}

	return nil
}
