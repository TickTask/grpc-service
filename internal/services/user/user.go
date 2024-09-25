package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"server/internal/domain/model"
	"server/internal/lib/jwt"
	"server/internal/storage"
	"time"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	log             *slog.Logger
	saverUser       SaverUser
	providerUser    ProviderUser
	sessionSaver    SessionSaver
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func New(
	log *slog.Logger,
	saverUser SaverUser,
	providerUser ProviderUser,
	sessionSaver SessionSaver,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) *User {
	return &User{
		log:             log,
		saverUser:       saverUser,
		providerUser:    providerUser,
		sessionSaver:    sessionSaver,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

type SaverUser interface {
	SaveUser(ctx context.Context, login string, passHash []byte, name string) (int64, error)
}

type ProviderUser interface {
	GetUser(ctx context.Context, login string) (model.User, error)
	GetUserByID(ctx context.Context, ID int64) (model.User, error)
}

type SessionSaver interface {
	SaveUserSession(ctx context.Context, userID int64, refreshToken string, sessionID string, deviceID string) error
	RefreshUserSession(ctx context.Context, deviceID string, userID int64, refreshToken string) error
}

func (u *User) Login(ctx context.Context, login string, password string, deviceID string) (model.Tokens, error) {
	const op = "user.login"

	var tokens model.Tokens

	log := u.log.With(slog.String("op", op), slog.String("login", login))

	log.Info("attempting to login")

	//Генерация ID сессии
	sessionID := uuid.NewString()

	//Получаем пользователя
	user, err := u.providerUser.GetUser(ctx, login)

	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			u.log.Warn("user not found", err.Error())
			return model.Tokens{}, errors.New("user not found")
		}

		u.log.Warn("error getting user", err.Error())
		return model.Tokens{}, errors.New("error getting user")
	}

	//Проверяем хэш паролей
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		u.log.Info("invalid password", err.Error())
		return model.Tokens{}, errors.New("invalid password")
	}

	u.log.Info("successfully logged in")

	//Генерация access токена
	access, err := jwt.NewAccessToken(user.ID, u.accessTokenTTL, sessionID, deviceID)

	if err != nil {
		u.log.Warn("error creating access token", err.Error())
		return model.Tokens{}, errors.New("error creating access token")
	}

	tokens.Access = access

	//Генерация refresh токена
	refresh, err := jwt.NewRefreshToken(user.ID, u.refreshTokenTTL, sessionID, deviceID)

	if err != nil {
		u.log.Warn("error creating refresh token", err.Error())
		return model.Tokens{}, errors.New("error creating refresh token")
	}

	tokens.Refresh = refresh

	//Добавление новой сессии пользователя

	if err := u.sessionSaver.SaveUserSession(ctx, user.ID, refresh, sessionID, deviceID); err != nil {
		u.log.Warn("error saving session", err.Error())
		return model.Tokens{}, errors.New("error saving session")
	}

	u.log.Info("successfully add session")

	return tokens, nil

}

func (u *User) Register(ctx context.Context, login string, password string, name string) (int64, error) {
	const op = "user.register"

	log := u.log.With(slog.String("op", op), slog.String("login", login))

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("Failed to hash password", err.Error())
		return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
	}

	id, err := u.saverUser.SaveUser(ctx, login, passHash, name)

	if err != nil {
		if errors.Is(err, ErrUserExists) {

			log.Warn("user already exists", err.Error())

			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		log.Error("Failed to save user", err.Error())

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (u *User) FetchUser(ctx context.Context, ID int64) (model.User, error) {
	const op = "user.fetch_user"

	log := u.log.With(slog.String("op", op))

	user, err := u.providerUser.GetUserByID(ctx, ID)

	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Error("user not found", err.Error())
			return model.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		log.Error("error getting user", err.Error())
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (u *User) RefreshToken(ctx context.Context, token string) (model.Tokens, error) {
	const op = "user.refresh_token"

	log := u.log.With(slog.String("op", op))

	var tokens model.Tokens

	t, err := jwt.ParseRefreshToken(token)

	if err != nil {
		return model.Tokens{}, fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err := jwt.NewAccessToken(t.UserID, u.accessTokenTTL, t.SessionID, t.DeviceID)

	if err != nil {
		log.Error("error creating access token", err.Error())
		return model.Tokens{}, errors.New("error creating access token")
	}

	refreshToken, err := jwt.NewRefreshToken(t.UserID, u.refreshTokenTTL, t.SessionID, t.DeviceID)

	if err != nil {
		log.Error("error creating access token", err.Error())
		return model.Tokens{}, errors.New("error creating access token")
	}

	err = u.sessionSaver.RefreshUserSession(ctx, t.DeviceID, t.UserID, refreshToken)

	if err != nil {
		log.Error("error saving refresh token", err.Error())
		return model.Tokens{}, errors.New("error saving refresh token")
	}
	tokens.Access = accessToken
	tokens.Refresh = refreshToken

	return tokens, nil
}
