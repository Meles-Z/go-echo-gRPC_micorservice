package interal

import (
	"github.com/labstack/echo/v4"
	productHandlers "github.com/meles-z/go-grpc-microsevice/interal/handlers/product"
	userHandlers "github.com/meles-z/go-grpc-microsevice/interal/handlers/user"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"google.golang.org/grpc"
)

// Routes registers all routes with the Echo server
func Routes(e *echo.Echo, grpcConn *grpc.ClientConn) {
	client := order.NewUserServiceClient(grpcConn)
	user := e.Group("/user")
	user.POST("/create", userHandlers.CreateUser(client))
	user.GET("/all", userHandlers.GetAllUsers(client))
	user.GET("/:id", userHandlers.GetUserById(client))
	user.PUT("/", userHandlers.UpdateUser(client))
	user.DELETE("/:id", userHandlers.DeleteUser(client))

	productClient := order.NewProductServiceClient(grpcConn)
	product := e.Group("/product")
	product.POST("/create", productHandlers.CreateProduct(productClient))
	product.GET("/all", productHandlers.GetAllProducts(productClient))
}
