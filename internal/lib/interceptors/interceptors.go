package interceptors

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"server/internal/lib/jwt"
	"strings"
)

var authFreeMethods = map[string]bool{
	"/user.User/RegisterUser": true,
	"/user.User/LoginUser":    true,
}

func IsAuth(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	if authFreeMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, fmt.Errorf("Error get metadata from context")
	}

	authHeader, ok := md["authorization"]

	if !ok || len(authHeader) == 0 {
		return nil, fmt.Errorf("Error get authorization header")
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

	if tokenString == authHeader[0] {
		return nil, fmt.Errorf("Error get authorization header")
	}

	claims, err := jwt.ParseAccessToken(tokenString)

	if err != nil {
		return nil, fmt.Errorf("Error parse token err: %w", err)
	}

	ctx = context.WithValue(ctx, "user_id", claims.UserID)

	ctx = context.WithValue(ctx, "session_id", claims.SessionID)

	ctx = context.WithValue(ctx, "device_id", claims.DeviceID)

	return handler(ctx, req)
}
