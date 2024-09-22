package application

import (
	"context"
	"time"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
)

type createOrderService struct {
	orderRepo     domain.OrderRepository
	orderLineRepo domain.OrderLineRepository
}

type CreateOrderInput struct {
	ID           string  `json:"id"`
	CustomerName string  `json:"customer_name"`
	Description  *string `json:"description"`
	OrderLines   []struct {
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
		ID:           s.orderRepo.NextID(),
		CustomerName: input.CustomerName,
		Description:  input.Description,
		CreatedOn:    time.Now().UTC(),
	}
	if err := s.orderRepo.Save(ctx, order); err != nil {
		return err
	}

	for _, line := range input.OrderLines {
		orderLine := &domain.OrderLine{
			ID:       s.orderLineRepo.NextID(),
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
