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
	GetAllUsers(ctx context.Context, req *order.GetAllUsersRequest) (*order.GetAllUsersResponse, error)
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

// Implement the GetAllUsers method
func (s *Server) GetAllUsers(ctx context.Context, req *order.GetAllUsersRequest) (*order.GetAllUsersResponse, error) {

	var dbUsers []entities.User
	if err := s.DB.Find(&dbUsers).Error; err != nil {
		return nil, errors.New("failed to return all users: " + err.Error())
	}

	// Convert database users to gRPC users
	var grpcUsers []*order.User
	for _, user := range dbUsers {
		grpcUsers = append(grpcUsers, &order.User{
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
			Phone:   user.Phone,
		})
	}

	return &order.GetAllUsersResponse{
		Users: grpcUsers,
	}, nil
}
