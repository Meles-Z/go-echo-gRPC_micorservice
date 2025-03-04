package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/interal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	e := echo.New()

	// Create gRPC connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Pass gRPC client connection to Routes
	interal.Routes(e, conn)

	// Start HTTP server
	port := 8080
	fmt.Printf("Server running on port %d\n", port)
	log.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
