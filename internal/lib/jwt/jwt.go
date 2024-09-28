package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/internal/domain/model"
	"time"
)

const (
	AccessTokenSigningKey  = "4Ap3YkO/3C2QEoSN29TYGYcFMZryjuM7jzdallEQI08="
	RefreshTokenSigningKey = "26H5WuJZGXAlVScP2UqY0iX/96wi8imTVXcABw+JIJQ="
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID    int64  `json:"user_id"`
	SessionID string `json:"session_id"`
	DeviceID  string `json:"device_id"`
}

func NewAccessToken(userID int64, duration time.Duration, sessionID string, deviceID string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID:    userID,
		SessionID: sessionID,
		DeviceID:  deviceID,
	})
	accessTokenString, err := accessToken.SignedString([]byte(AccessTokenSigningKey))
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}

func NewRefreshToken(userID int64, duration time.Duration, sessionID string, deviceID string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID:    userID,
		SessionID: sessionID,
		DeviceID:  deviceID,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(RefreshTokenSigningKey))
	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}

func ParseRefreshToken(requestToken string) (model.ParseTokens, error) {
	var parsedToken model.ParseTokens
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(RefreshTokenSigningKey), nil
	})

	if err != nil {
		fmt.Println("Ошибка парсинга токена:", err)
		return model.ParseTokens{}, err
	}

	if !token.Valid {
		return model.ParseTokens{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return model.ParseTokens{}, errors.New("Token claims are not of type *refreshTokenClaims")
	}

	parsedToken.UserID = claims.UserID
	parsedToken.SessionID = claims.SessionID
	parsedToken.DeviceID = claims.DeviceID

	return parsedToken, nil
}

func ParseAccessToken(requestToken string) (model.ParseTokens, error) {
	var parsedToken model.ParseTokens
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AccessTokenSigningKey), nil
	})
	if err != nil {
		fmt.Println("Ошибка парсинга токена:", err)
		return model.ParseTokens{}, nil
	}
	if !token.Valid {
		return model.ParseTokens{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return model.ParseTokens{}, errors.New("Token claims are not of type *refreshTokenClaims")
	}

	parsedToken.UserID = claims.UserID
	parsedToken.SessionID = claims.SessionID
	parsedToken.DeviceID = claims.DeviceID

	return parsedToken, nil
}
