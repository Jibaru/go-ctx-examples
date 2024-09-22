package domain

type Order struct {
	ID         string `json:"id" bson:"_id"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
}

type OrderLine struct {
	ID       string `json:"id" bson:"_id"`
	OrderID  string `json:"order_id" bson:"order_id"`
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
