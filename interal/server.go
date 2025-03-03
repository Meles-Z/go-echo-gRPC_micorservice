package interal

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/meles-z/go-grpc-microsevice/config"
	"github.com/meles-z/go-grpc-microsevice/interal/database"
	"github.com/meles-z/go-grpc-microsevice/interal/repository"
	"github.com/meles-z/go-grpc-microsevice/interal/service"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type IServer interface {
	Start() error
}

type Server struct {
	DB     *gorm.DB
	config config.Config
	user   service.IUserService
}

func NewServer(cfg config.Config) IServer {
	// Initialize database
	mainDb, err := database.InitDB(&cfg.DB)
	if err != nil {
		log.Fatalf("DB INITIALIZE ERROR :%s", err.Error())
	}

	// Initialize user repository and service
	userRepository := repository.NewUserRepository(mainDb)
	userSvc, err := service.NewUserService(userRepository)
	if err != nil {
		log.Fatal("Error to initialize user service:", err)
	}

	// Return the server instance
	return &Server{
		DB:     mainDb,
		config: cfg,
		user:   userSvc,
	}
}

func (s *Server) Start() error {
	fmt.Println("gRPC running...")

	// Listen on the specified port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the user service with the gRPC server
	order.RegisterUserServiceServer(grpcServer, &repository.Server{
		DB: s.DB,
		UnimplementedUserServiceServer: &order.UnimplementedUserServiceServer{},
	})

	log.Printf("Server listening at %v", lis.Addr())

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
