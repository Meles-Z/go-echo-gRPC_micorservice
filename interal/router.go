package interal

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/meles-z/go-grpc-microsevice/interal/handlers/user"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"google.golang.org/grpc"
)

// Routes registers all routes with the Echo server
func Routes(e *echo.Echo, grpcConn *grpc.ClientConn) {
	client := order.NewUserServiceClient(grpcConn)
	user := e.Group("/user")
	user.POST("/create", handlers.CreateUser(client))
	user.GET("/all", handlers.GetAllUsers(client))
	user.GET("/get-user/:id", handlers.GetUserById(client))
}
