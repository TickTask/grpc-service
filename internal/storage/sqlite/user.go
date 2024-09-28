package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	sqlite3 "github.com/mutecomm/go-sqlcipher/v4"
	"server/internal/domain/model"
	"server/internal/storage"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.new"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

func (s *Storage) SaveUser(ctx context.Context, login string, passHash []byte, name string) (int64, error) {
	const op = "storage.sqlite.save_user"

	req, err := s.db.Prepare("INSERT INTO Users(login, name, hash_password) VALUES (?, ?, ?)")

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := req.ExecContext(ctx, login, name, passHash)

	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExist)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetUser(ctx context.Context, login string) (model.User, error) {
	const op = "storage.sqlite.get_user"

	req, err := s.db.Prepare("SELECT id, login, name, hash_password FROM Users WHERE login = ?")
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}
	row := req.QueryRowContext(ctx, login)

	var user model.User

	err = row.Scan(&user.ID, &user.Login, &user.Name, &user.PassHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) GetUserByID(ctx context.Context, ID int64) (model.User, error) {
	const op = "storage.sqlite.get_user_by_id"

	var user model.User

	req, err := s.db.Prepare("SELECT id, login, name FROM Users WHERE id = ?")

	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := req.QueryRowContext(ctx, ID)

	err = row.Scan(&user.ID, &user.Login, &user.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) SaveUserSession(ctx context.Context, userID int64, refreshToken string, sessionID string, deviceID string) error {
	const op = "storage.sqlite.save_user_session"

	req, err := s.db.Prepare("INSERT INTO Sessions(id, refresh_token, session_user_id, device_id) VALUES (?, ?, ?, ?)")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = req.ExecContext(ctx, sessionID, refreshToken, userID, deviceID)

	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {

			return fmt.Errorf("%s: %w", op, storage.ErrSessionExist)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) RefreshUserSession(ctx context.Context, deviceID string, userID int64, newToken string, sessionID string, oldToken string) error {
	const op = "storage.sqlite.refresh_user_session"

	req, err := s.db.Prepare("UPDATE Sessions SET refresh_token = ? WHERE device_id = ? AND session_user_id = ? AND id = ? AND refresh_token = ?")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	r, err := req.ExecContext(ctx, newToken, deviceID, userID, sessionID, oldToken)

	rows, err := r.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrSessionNotFound)
	}

	return nil
}

func (s *Storage) RemoveUserSession(ctx context.Context, sessionID string, userID int64, deviceID string) error {
	const op = "storage.sqlite.remove_user_session"

	req, err := s.db.Prepare("DELETE FROM Sessions WHERE session_user_id = ? AND id = ?  AND device_id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = req.ExecContext(ctx, userID, sessionID, deviceID)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
