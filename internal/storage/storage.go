package storage

import "errors"

var (
	ErrUserExist = errors.New("user already exist")

	ErrSessionExist = errors.New("session already exist")

	ErrSessionNotFound = errors.New("session not found")

	ErrUserNotFound = errors.New("user not found")

	ErrTaskNotFound = errors.New("task not found")
)
