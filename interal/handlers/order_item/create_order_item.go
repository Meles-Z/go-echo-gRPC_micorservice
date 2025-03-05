package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

type CustomError struct {
	Reason string `json:"reason"`
}

func CreateOrderItem(client order.OrderItemServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderItem entities.OrderItem
		if err := c.Bind(&orderItem); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, &CustomError{
				Reason: fmt.Sprintf("invalid payload: %s", err.Error()),
			})
		}
		newItem, err := client.CreateOrderItem(context.Background(), &order.CreateOrderItemRequest{
			OrderItem: &order.OrderItem{
				OrderId:    orderItem.OrderID,
				ProductId:  orderItem.ProductID,
				Quantity:   int32(orderItem.Quantity),
				UnitPrice:  orderItem.UnitPrice,
				TotalPrice: orderItem.TotalPrice,
			},
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &CustomError{
				Reason: fmt.Sprintf("error to creating order item:%s", err.Error()),
			})
		}
		res := map[string]any{
			"orderItem:": newItem.OrderItem,
		}
		return c.JSON(http.StatusOK, res)
	}
}
