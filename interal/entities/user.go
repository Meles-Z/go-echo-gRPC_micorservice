package entities

type User struct {
	Model
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
	Password string  `json:"-"`
	Orders   []Order `json:"orders"` // One-to-many relationship with orders
}
