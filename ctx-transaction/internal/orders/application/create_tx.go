package application

import (
	"context"
)

type createOrderServiceTx struct {
	createOrderService CreateOrderService
	transactional      Transactional
}

func NewCreateOrderServiceTx(
	createOrderService CreateOrderService,
	transactional Transactional,
) *createOrderServiceTx {
	return &createOrderServiceTx{
		createOrderService: createOrderService,
		transactional:      transactional,
	}
}

func (s *createOrderServiceTx) Exec(ctx context.Context, input CreateOrderInput) error {
	return s.transactional.InTransaction(ctx, func(txCtx context.Context) error {
		return s.createOrderService.Exec(txCtx, input)
	})
}
