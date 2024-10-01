package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"server/internal/domain/model"
	"server/internal/storage"
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
	const op = "storage.sqlite.get_task_by_id"

	var task model.Task
	var user model.TodosUser
	var status model.Status

	req, err := t.db.Prepare(`SELECT t.id, t.title, t.body, t.created_at, u.id, u.name, s.id, s.status FROM Tasks t 
    INNER JOIN Users u ON u.id = t.task_user_id 
    INNER JOIN Statuses s ON t.task_status_id = s.id WHERE t.id = ?`)

	defer req.Close()

	if err != nil {
		return task, fmt.Errorf("%s: %w", op, err)
	}

	row := req.QueryRowContext(ctx, taskID)

	err = row.Scan(&task.ID, &task.Title, &task.Body, &task.CreatedAt, &user.ID, &user.Name, &status.ID, &status.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, fmt.Errorf("%s: %w", op, storage.ErrTaskNotFound)
		}
		return model.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	task.User = user

	task.Status = status

	return task, nil
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
