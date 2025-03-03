package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-grpc-microsevice/config"
	"github.com/meles-z/go-grpc-microsevice/interal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error to load env file:", err)
	}
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error to load config file:", err)
	}
	db, err := database.InitDB(&config.DB)
	if err != nil {
		log.Fatalln("Error to initateDb:", err)
	}
	fmt.Print("db:", db)

	ech := echo.New()
	log.Fatal(ech.Start(fmt.Sprintf(":%d", config.Server.Port)))
}
