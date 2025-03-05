package repository

import (
	"context"
	"errors"
	"fmt"

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
	GetUserById(ctx context.Context, req *order.GetUserByIdRequest) (*order.GetUserByIdResponse, error)
	UpdateUser(cxt context.Context, req *order.UpdateUserRequest) (*order.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *order.DeleteUserRequest) (*order.DeleteUserResponse, error)
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

func (s *Server) GetUserById(ctx context.Context, req *order.GetUserByIdRequest) (*order.GetUserByIdResponse, error) {
	var user entities.User
	if err := s.DB.Where("id=?", req.GetId()).Take(&user).Error; err != nil {
		return nil, errors.New("User not found: " + err.Error())
	}
	return &order.GetUserByIdResponse{
		User: &order.User{
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
			Phone:   user.Phone,
		},
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *order.UpdateUserRequest) (*order.UpdateUserResponse, error) {
	var existingUser entities.User
	reqUser := req.GetUser()
	fmt.Println("User id response:", reqUser.Id)
	err := s.DB.Model(&existingUser).Where("id=?", reqUser.Id).Updates(entities.User{
		Name:     reqUser.Name,
		Email:    reqUser.Email,
		Address:  reqUser.Address,
		Phone:    reqUser.Phone,
		Password: reqUser.Password,
	}).Scan(&existingUser).Error
	if err != nil {
		return nil, errors.New("failed to update users:" + err.Error())
	}
	return &order.UpdateUserResponse{
		User: &order.User{
			Id:       existingUser.ID,
			Name:     existingUser.Name,
			Email:    existingUser.Email,
			Address:  existingUser.Address,
			Phone:    existingUser.Phone,
			Password: existingUser.Password,
		},
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *order.DeleteUserRequest) (*order.DeleteUserResponse, error) {
	var user entities.User
	result := s.DB.Where("id=?", req.GetId()).Delete(&user) // Get the result

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) { // Check for "record not found"
			return nil, errors.New("user not found")
		}
		return nil, errors.New("Error to delete user: " + result.Error.Error())
	}

	if result.RowsAffected == 0 { // Check if any rows were affected
		return nil, errors.New("user not found")
	}

	return &order.DeleteUserResponse{
		Success: true,
	}, nil
}
