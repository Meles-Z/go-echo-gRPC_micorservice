package entities

type Order struct {
	Model
	UserId     string      `json:"userId"`
	ProductId  string      `json:"productId"`
	Price      int         `json:"price"`
	Quantity   int         `json:"quantity"`
	Status     string      `json:"status"`
	User       User        `gorm:"foreignKey:UserID"` // GORM Foreign Key relationship with User
	OrderItems []OrderItem `json:"order_items"`       // Relationship: One-to-Many with OrderItems
}
