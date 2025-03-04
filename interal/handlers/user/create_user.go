package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

// CreateUser handles user creation via gRPC
func CreateUser(client order.UserServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entities.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("error binding user: %v", err))
		}

		// Prepare gRPC request

		newUser := &order.User{
			Name:     user.Name,
			Email:    user.Email,
			Address:  user.Address,
			Phone:    user.Phone,
			Password: user.Password,
		}

		// Call gRPC server
		res, err := client.CreateUser(context.Background(), &order.CreateUserRequest{
			User: newUser,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating user: %v", err))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "User created successfully",
			"data":    res.User,
		})
	}
}