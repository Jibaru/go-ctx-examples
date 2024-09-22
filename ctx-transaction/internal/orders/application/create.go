package application

import (
	"context"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
)

type createOrderService struct {
	orderRepo     domain.OrderRepository
	orderLineRepo domain.OrderLineRepository
}

type CreateOrderInput struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	OrderLines []struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	} `json:"order_lines"`
}

func NewCreateOrderService(
	orderRepo domain.OrderRepository,
	orderLineRepo domain.OrderLineRepository,
) *createOrderService {
	return &createOrderService{
		orderRepo:     orderRepo,
		orderLineRepo: orderLineRepo,
	}
}

func (s *createOrderService) Exec(ctx context.Context, input CreateOrderInput) error {
	order := &domain.Order{
		ID:         "1",
		CustomerID: input.CustomerID,
	}
	if err := s.orderRepo.Save(ctx, order); err != nil {
		return err
	}

	for _, line := range input.OrderLines {
		orderLine := &domain.OrderLine{
			OrderID:  order.ID,
			Name:     line.Name,
			Quantity: line.Quantity,
		}

		if err := s.orderLineRepo.Save(ctx, orderLine); err != nil {
			return err
		}
	}
	return nil
}
