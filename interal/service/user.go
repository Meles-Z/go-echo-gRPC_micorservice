package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/meles-z/go-grpc-microsevice/interal/repository"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *order.CreateUserRequest) (*order.CreateUserResponse, error)
	GetAllUsers(ctx context.Context, req *order.GetAllUsersRequest) (*order.GetAllUsersResponse, error)
	GetUserById(ctx context.Context, req *order.GetUserByIdRequest) (*order.GetUserByIdResponse, error)
	UpdateUser(ctx context.Context, req *order.UpdateUserRequest) (*order.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *order.DeleteUserRequest) (*order.DeleteUserResponse, error)
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

func (srv *UserService) GetAllUsers(ctx context.Context, req *order.GetAllUsersRequest) (*order.GetAllUsersResponse, error) {
	users, err := srv.userRepo.GetAllUsers(ctx, req)
	if err != nil {
		return nil, errors.New("failed to feach all users: " + err.Error())
	}
	return users, nil
}

func (srv *UserService) GetUserById(ctx context.Context, req *order.GetUserByIdRequest) (*order.GetUserByIdResponse, error) {
	user, err := srv.userRepo.GetUserById(ctx, req)
	if err != nil {
		return nil, errors.New("failed to feach users by id: " + err.Error())
	}
	return user, nil
}

func (svc *UserService) UpdateUser(ctx context.Context, req *order.UpdateUserRequest) (*order.UpdateUserResponse, error) {
	fmt.Println("I am here in service")
	fmt.Println("userid:", req.User.Id)

	updatedUser, err := svc.userRepo.UpdateUser(ctx, req)
	if err != nil {
		return nil, errors.New("failed to update user:" + err.Error())
	}
	return updatedUser, nil
}

func (svc *UserService) DeleteUser(ctx context.Context, req *order.DeleteUserRequest) (*order.DeleteUserResponse, error) {

	deletedUser, err := svc.userRepo.DeleteUser(ctx, req)
	if err != nil {
		return nil, errors.New("error to delete user:" + err.Error())
	}
	return deletedUser, nil
}