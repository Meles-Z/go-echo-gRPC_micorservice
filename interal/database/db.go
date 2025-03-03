package database

import (
	"fmt"

	"github.com/meles-z/go-grpc-microsevice/config"
	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Name, cfg.Username, cfg.Password)
	orderDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = orderDb
	orderDb.AutoMigrate(
		&entities.User{},
		&entities.Product{},
		&entities.OrderItem{},
		&entities.Order{},
	)
	fmt.Println("Databas connected successfully!")
	return orderDb, nil
}
