package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

type ErrorReason struct {
	Reason string `json:"reason"`
}

func UpdateUser(client order.UserServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entities.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, &ErrorReason{
				Reason: "unformated payload" + err.Error(),
			})
		}
		res, err := client.UpdateUser(context.Background(), &order.UpdateUserRequest{
			User: &order.User{
				Id:       user.ID,
				Name:     user.Name,
				Email:    user.Email,
				Address:  user.Address,
				Phone:    user.Phone,
				Password: user.Password,
			},
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &ErrorReason{
				Reason: fmt.Sprintf("Error to update user" + err.Error()),
			})
		}
		respp := map[string]any{
			"updatedUser:": res.User,
		}
		return c.JSON(http.StatusOK, respp)
	}
}
