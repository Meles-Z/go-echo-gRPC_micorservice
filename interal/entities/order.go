package entities

type Order struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	ProductId string `json:"productId"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}
