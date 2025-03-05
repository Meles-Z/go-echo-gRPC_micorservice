package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

func DeleteUser(client order.UserServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		deletedUser, err := client.DeleteUser(context.Background(), &order.DeleteUserRequest{
			Id: id,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error:": err.Error(),
			})
		}
		resp := map[string]interface{}{
			"success": deletedUser.Success,
		}
		return c.JSON(http.StatusOK, resp)
	}
}
