package entities

type Product struct {
	Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Orders      []Order `json:"orders" gorm:"many2many:order_items;"` // Corrected Many-to-Many Relationship
}
