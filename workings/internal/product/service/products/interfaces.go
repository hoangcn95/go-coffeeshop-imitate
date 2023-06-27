package products

import (
	"context"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/domain"
)

type ProductService interface {
	GetItemTypes(context.Context) ([]*domain.ItemTypeDto, error)
	GetItemsByType(context.Context, string) ([]*domain.ItemDto, error)
}
