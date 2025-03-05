package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

func GetAllUsers(client order.UserServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Call gRPC service
		res, err := client.GetAllUsers(context.Background(), &order.GetAllUsersRequest{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("failed to fetch users: %v", err),
			})
		}
		return c.JSON(http.StatusOK, res)
	}
}
