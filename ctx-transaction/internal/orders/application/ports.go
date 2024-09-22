package application

import (
	"context"
)

type Transactional interface {
	InTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type CreateOrderService interface {
	Exec(ctx context.Context, input CreateOrderInput) error
}
