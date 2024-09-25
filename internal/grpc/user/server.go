package user

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/internal/domain/model"
	"server/internal/services/user"
	userRpc "server/pkg/user"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type User interface {
	Login(ctx context.Context, login string, password string, deviceID string) (model.Tokens, error)
	Register(ctx context.Context, login string, password string, name string) (int64, error)
	FetchUser(ctx context.Context, ID int64) (model.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (model.Tokens, error)
}

// Хэндлеры
type serverApi struct {
	userRpc.UnimplementedUserServer
	user User
}

func Register(gRPC *grpc.Server, user User) {
	userRpc.RegisterUserServer(gRPC, &serverApi{user: user})
}

func (s *serverApi) LoginUser(ctx context.Context, req *userRpc.LoginUserRequest) (*userRpc.LoginUserResponse, error) {
	err := validateLogin(req)

	if err != nil {
		return nil, err
	}

	tokens, err := s.user.Login(ctx, req.GetLogin(), req.GetPassword(), req.GetDeviceId())

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid credentials")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &userRpc.LoginUserResponse{
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	}, nil
}

func (s *serverApi) RegisterUser(ctx context.Context, request *userRpc.RegisterUserRequest) (*userRpc.RegisterUserResponse, error) {
	if err := validateRegister(request); err != nil {
		return nil, err
	}

	id, err := s.user.Register(ctx, request.GetLogin(), request.GetPassword(), request.GetUsername())
	if err != nil {
		if errors.Is(err, ErrUserExists) {
			return nil, status.Error(codes.InvalidArgument, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &userRpc.RegisterUserResponse{UserId: id}, nil
}

func (s *serverApi) GetUser(ctx context.Context, request *userRpc.GetUserRequest) (*userRpc.GetUserResponse, error) {
	err := validateFetchUser(request)
	if err != nil {
		return nil, err
	}
	u, err := s.user.FetchUser(ctx, request.GetUserId())
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &userRpc.GetUserResponse{
		UserId:   u.ID,
		Username: u.Name,
		Login:    u.Login,
	}, nil
}

func (s *serverApi) RefreshToken(ctx context.Context, request *userRpc.RefreshTokenRequest) (*userRpc.RegisterUserResponse, error) {
	err := validateRefreshToken(request)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func validateLogin(req *userRpc.LoginUserRequest) error {
	if req.GetLogin() == "" {
		return status.Error(codes.InvalidArgument, "Login is required")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "Password is required")
	}
	return nil
}

func validateRegister(req *userRpc.RegisterUserRequest) error {
	if req.GetLogin() == "" {
		return status.Error(codes.InvalidArgument, "Login is required")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "Password is required")
	}

	if req.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}

	return nil
}

func validateFetchUser(req *userRpc.GetUserRequest) error {
	if req.GetUserId() == 0 || req == nil {
		return status.Error(codes.InvalidArgument, "UserId is required")
	}
	return nil
}

func validateRefreshToken(req *userRpc.RefreshTokenRequest) error {
	if req.GetRefreshToken() == "" {
		return status.Error(codes.InvalidArgument, "RefreshToken is required")
	}
	return nil
}
