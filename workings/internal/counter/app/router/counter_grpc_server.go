package router

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/cmd/counter/config"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/domain"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/service/order"
	shared "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/shared_kernel"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/proto/gen"
)

type counterGRPCServer struct {
	gen.UnimplementedCounterServiceServer
	cfg *config.Config
	os  order.OrderService
}

var _ gen.CounterServiceServer = (*counterGRPCServer)(nil)

var CounterGRPCServerSet = wire.NewSet(NewGRPCCounterServer)

func NewGRPCCounterServer(
	grpcServer *grpc.Server,
	cfg *config.Config,
	os order.OrderService,
) gen.CounterServiceServer {
	svc := counterGRPCServer{
		cfg: cfg,
		os:  os,
	}

	gen.RegisterCounterServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *counterGRPCServer) GetListOrderFulfillment(
	ctx context.Context,
	request *gen.GetListOrderFulfillmentRequest,
) (*gen.GetListOrderFulfillmentResponse, error) {
	slog.Info("GET: GetListOrderFulfillment")

	res := gen.GetListOrderFulfillmentResponse{}

	entities, err := g.os.GetListOrderFulfillment(ctx)
	if err != nil {
		return nil, fmt.Errorf("uc.GetListOrderFulfillment: %w", err)
	}

	for _, entity := range entities {
		res.Orders = append(res.Orders, &gen.OrderDto{
			Id:              entity.ID.String(),
			OrderSource:     int32(entity.OrderSource),
			OrderStatus:     int32(entity.OrderStatus),
			Location:        int32(entity.Location),
			LoyaltyMemberId: entity.LoyaltyMemberID.String(),
			LineItems: lo.Map(entity.LineItems, func(item *domain.LineItem, _ int) *gen.LineItemDto {
				return &gen.LineItemDto{
					Id:             item.ID.String(),
					ItemType:       int32(item.ItemType),
					Name:           item.Name,
					Price:          float64(item.Price),
					ItemStatus:     int32(item.ItemStatus),
					IsBaristaOrder: item.IsBaristaOrder,
				}
			}),
		})
	}

	return &res, nil
}

func (g *counterGRPCServer) PlaceOrder(
	ctx context.Context,
	request *gen.PlaceOrderRequest,
) (*gen.PlaceOrderResponse, error) {
	slog.Info("POST: PlaceOrder")

	loyaltyMemberID, err := uuid.Parse(request.LoyaltyMemberId)
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	model := domain.PlaceOrderModel{
		CommandType:     shared.CommandType(request.CommandType),
		OrderSource:     shared.OrderSource(request.OrderSource),
		Location:        shared.Location(request.Location),
		LoyaltyMemberID: loyaltyMemberID,
		Timestamp:       request.Timestamp.AsTime(),
	}

	for _, barista := range request.BaristaItems {
		model.BaristaItems = append(model.BaristaItems, &domain.OrderItemModel{
			ItemType: shared.ItemType(barista.ItemType),
		})
	}

	for _, kitchen := range request.KitchenItems {
		model.KitchenItems = append(model.KitchenItems, &domain.OrderItemModel{
			ItemType: shared.ItemType(kitchen.ItemType),
		})
	}

	err = g.os.PlaceOrder(ctx, &model)
	if err != nil {
		return nil, errors.Wrap(err, "uc.PlaceOrder")
	}

	res := gen.PlaceOrderResponse{}

	return &res, nil
}
