package entities

type OrderItem struct {
	ID         uint    `json:"id"`
	OrderID    uint    `json:"order_id"`   // Foreign key to Order
	ProductID  uint    `json:"product_id"` // Foreign key to Product
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"` // Quantity * UnitPrice
	Product    Product `json:"product"`
	Order      Order   `json:"order"`
}
