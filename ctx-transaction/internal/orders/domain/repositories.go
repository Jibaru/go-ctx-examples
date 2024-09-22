package domain

import "context"

type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
}

type OrderLineRepository interface {
	Save(ctx context.Context, orderLine *OrderLine) error
}
