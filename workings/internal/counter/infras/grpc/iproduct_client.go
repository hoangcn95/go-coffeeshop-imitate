package grpc

import (
	"context"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/domain"
)

type (
	ProductDomainService interface {
		GetItemsByType(context.Context, *domain.PlaceOrderModel, bool) ([]*domain.ItemModel, error)
	}
)
