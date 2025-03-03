package entities

import "time"

type Product struct {
	Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	StockQty    int         `json:"stock_qty"`
	CreatedAt   time.Time   `json:"created_at"`
	Orders      []OrderItem `json:"orders" gorm:"foreignKey:many2many"` // Many-to-many relationship with orders through OrderItems
}
