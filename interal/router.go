package interal

import (
	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/interal/handlers"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"google.golang.org/grpc"
)

// Routes registers all routes with the Echo server
func Routes(e *echo.Echo, grpcConn *grpc.ClientConn) {
	client := order.NewUserServiceClient(grpcConn)
	e.POST("/users", handlers.CreateUser(client))
}
