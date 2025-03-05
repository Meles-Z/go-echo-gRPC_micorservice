package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
)

func GetAllProducts(client order.ProductServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		product, err := client.GetAllProducts(context.Background(), &order.GetAllProductsRequest{})
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"record not found": err.Error(),
			})
		}
		resp := map[string]any{
			"products:": product.Product,
		}
		return c.JSON(http.StatusOK, resp)

	}
}
