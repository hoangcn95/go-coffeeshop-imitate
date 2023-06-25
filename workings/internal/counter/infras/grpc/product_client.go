package grpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/cmd/counter/config"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/domain"
	shared "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/shared_kernel"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/proto/gen"
)

type productGRPCClient struct {
	conn *grpc.ClientConn
}

var _ ProductDomainService = (*productGRPCClient)(nil)

// definition custom dependency injection stuff
var ProductGRPCClientSet = wire.NewSet(NewGRPCProductClient)

func NewGRPCProductClient(cfg *config.Config) (ProductDomainService, error) {
	conn, err := grpc.Dial(cfg.ProductClient.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &productGRPCClient{
		conn: conn,
	}, nil
}

func (p *productGRPCClient) GetItemsByType(
	ctx context.Context,
	model *domain.PlaceOrderModel,
	isBarista bool,
) ([]*domain.ItemModel, error) {
	c := gen.NewProductServiceClient(p.conn)

	itemTypes := ""
	if isBarista {
		itemTypes = lo.Reduce(model.BaristaItems, func(agg string, item *domain.OrderItemModel, _ int) string {
			return fmt.Sprintf("%s,%s", agg, item.ItemType.String())
		}, "")
	} else {
		itemTypes = lo.Reduce(model.KitchenItems, func(agg string, item *domain.OrderItemModel, _ int) string {
			return fmt.Sprintf("%s,%s", agg, item.ItemType.String())
		}, "")
	}

	req := &gen.GetItemsByTypeRequest{
		ItemTypes: strings.TrimLeft(itemTypes, ","),
	}
	res, err := c.GetItemsByType(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "productGRPCClient-c.GetItemsByType")
	}

	results := make([]*domain.ItemModel, 0)
	for _, item := range res.Items {
		results = append(results, &domain.ItemModel{
			ItemType: shared.ItemType(item.Type),
			Price:    item.Price,
		})
	}

	return results, nil
}
