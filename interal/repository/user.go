package repository

import (
	"context"
	"errors"

	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	*order.UnimplementedUserServiceServer
}
type UserRepository interface {
	CreateUser(ctx context.Context, req *order.CreateUserRequest) (*order.CreateUserResponse, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &Server{
		DB: db,
	}
}

func (s *Server) CreateUser(ctx context.Context, req *order.CreateUserRequest) (*order.CreateUserResponse, error) {
	user := req.GetUser()
	data := entities.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: user.Password,
	}
	res := s.DB.Create(&data)
	if res.Error != nil {
		return nil, errors.New("failed to create user: " + res.Error.Error())
	}
	return &order.CreateUserResponse{
		User: &order.User{
			Name:     user.Name,
			Email:    user.Email,
			Address:  user.Address,
			Phone:    user.Phone,
			Password: user.Password,
		},
	}, nil
}
