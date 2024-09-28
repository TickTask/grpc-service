package model

import "time"

type RequestTask struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64     `json:"user_id"`
	StatusID  int64     `json:"status_id"`
}

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UserID    TodosUser `json:"user"`
	StatusID  Status    `json:"status"`
}
