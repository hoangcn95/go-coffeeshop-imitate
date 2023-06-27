package products

import (
	"context"
	"strings"

	"github.com/google/wire"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/domain"
	"github.com/pkg/errors"
)

type productService struct {
	repo domain.ProductRepo
}

var _ ProductService = (*productService)(nil)

// definition a custom dependency injection
var ProductServiceSet = wire.NewSet(NewProductService)

func NewProductService(repo domain.ProductRepo) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetItemTypes(ctx context.Context) ([]*domain.ItemTypeDto, error) {
	results, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}

	return results, nil
}

func (s *productService) GetItemsByType(ctx context.Context, itemTypes string) ([]*domain.ItemDto, error) {
	types := strings.Split(itemTypes, ",")

	results, err := s.repo.GetByTypes(ctx, types)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemsByType")
	}

	return results, nil
}
