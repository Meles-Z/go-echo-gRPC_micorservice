package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

func CreateProduct(client order.ProductServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entities.Product
		if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]any{
				"failed to bind payload:": err.Error(),
			})
		}
		prod, err := client.CreateProduct(context.Background(), &order.CreateProductRequest{
			Product: &order.Product{
				Name:          product.Name,
				Description:   product.Description,
				Price:         product.Price,
				StockQuantity: int64(product.StockQty),
			},
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"invalid payload": err.Error(),
			})
		}
		res := map[string]any{
			"product": prod.Product,
		}
		return c.JSON(http.StatusOK, res)
	}
}
