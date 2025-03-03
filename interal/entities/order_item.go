package entities

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"` // Foreign key to Order
	ProductID  uint    `gorm:"not null"` // Foreign key to Product
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`          // Quantity * UnitPrice
	Order      Order   `gorm:"foreignKey:OrderID"`   // GORM Foreign Key relationship with Order
	Product    Product `gorm:"foreignKey:ProductID"` // GORM Foreign Key relationship with Product
}
