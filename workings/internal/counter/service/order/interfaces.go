package order

import (
	"context"

	"github.com/google/uuid"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/domain"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq/publisher"
)

type (
	OrderRepo interface {
		GetAll(context.Context) ([]*domain.Order, error)
		GetByID(context.Context, uuid.UUID) (*domain.Order, error)
		Create(context.Context, *domain.Order) error
		Update(context.Context, *domain.Order) (*domain.Order, error)
	}

	BaristaEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}

	KitchenEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}

	OrderService interface {
		GetListOrderFulfillment(context.Context) ([]*domain.Order, error)
		PlaceOrder(context.Context, *domain.PlaceOrderModel) error
	}
)
