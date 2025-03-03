package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/meles-z/go-grpc-microsevice/config"
	"github.com/meles-z/go-grpc-microsevice/interal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not load .env file", err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Could not load config file.", err)
	}
	inter := interal.NewServer(*cfg)
	log.Fatalln(inter.Start())
}
