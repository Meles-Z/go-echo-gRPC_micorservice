package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

func GetUserById(client order.UserServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		user, err := client.GetUserById(context.Background(), &order.GetUserByIdRequest{
			Id: id,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error: ": err.Error(),
			})
		}
		resp := map[string]any{
			"user": user.User,
		}
		return c.JSON(http.StatusOK, resp)
	}
}
