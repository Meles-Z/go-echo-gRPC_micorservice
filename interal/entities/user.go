package entities

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name"`
    Email     string    `json:"email" gorm:"unique"`
    Address   string    `json:"address"`
    Phone     string    `json:"phone"`
    Password  string    `json:"-"` // Do not expose password
    Orders    []Order   `json:"orders"` // One-to-many relationship with orders
}
