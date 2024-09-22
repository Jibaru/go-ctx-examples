package domain

import "time"

type Order struct {
	ID           any       `json:"id" bson:"_id"`
	CustomerName string    `json:"customer_name" bson:"customer_name"`
	Description  *string   `json:"description" bson:"description"`
	CreatedOn    time.Time `json:"created_on" bson:"created_on"`
}

type OrderLine struct {
	ID       any    `json:"id" bson:"_id"`
	OrderID  any    `json:"order_id" bson:"order_id"`
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
