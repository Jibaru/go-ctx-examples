package domain

import "context"

type OrderRepository interface {
	NextID() any
	Save(ctx context.Context, order *Order) error
}

type OrderLineRepository interface {
	NextID() any
	Save(ctx context.Context, orderLine *OrderLine) error
}
