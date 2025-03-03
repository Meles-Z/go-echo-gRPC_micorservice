package service

import (
	"context"
	"errors"

	"github.com/meles-z/go-grpc-microsevice/interal/repository"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *order.CreateUserRequest) (*order.CreateUserResponse, error)
}

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) (IUserService, error) {
	return &UserService{
		userRepo: userRepo,
	}, nil
}

func (srv *UserService) CreateUser(ctx context.Context, req *order.CreateUserRequest) (*order.CreateUserResponse, error) {
	user, err := srv.userRepo.CreateUser(ctx, req)
	if err != nil {
		// You could wrap the error for more context if needed
		return nil, errors.New("failed to create user: " + err.Error())
	}
	return user, nil
}
