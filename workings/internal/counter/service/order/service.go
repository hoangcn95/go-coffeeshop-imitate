package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/domain"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/infras/grpc"
	events "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/events"
	shared "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/shared_kernel"
)

type orderService struct {
	orderRepo        OrderRepo
	productDomainSvc grpc.ProductDomainService
	baristaEventPub  BaristaEventPublisher
	kitchenEventPub  KitchenEventPublisher
}

var _ OrderService = (*orderService)(nil)

var OrderServiceSet = wire.NewSet(NewOrderService)

func NewOrderService(
	orderRepo OrderRepo,
	productDomainSvc grpc.ProductDomainService,
	baristaEventPub BaristaEventPublisher,
	kitchenEventPub KitchenEventPublisher,
) OrderService {
	return &orderService{
		orderRepo:        orderRepo,
		productDomainSvc: productDomainSvc,
		baristaEventPub:  baristaEventPub,
		kitchenEventPub:  kitchenEventPub,
	}
}

func (os *orderService) GetListOrderFulfillment(ctx context.Context) ([]*domain.Order, error) {
	entities, err := os.orderRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("orderRepo.GetAll: %w", err)
	}

	return entities, nil
}

func (os *orderService) PlaceOrder(ctx context.Context, model *domain.PlaceOrderModel) error {
	order, err := os.createOrderFrom(ctx, model)
	if err != nil {
		return errors.Wrap(err, "domain.CreateOrderFrom")
	}

	err = os.orderRepo.Create(ctx, order)
	if err != nil {
		return errors.Wrap(err, "orderRepo.Create")
	}

	slog.Debug("order created", "order", *order)

	// todo: it might cause dual-write problem, but we accept it temporary
	for _, event := range order.DomainEvents() {
		if event.Identity() == "BaristaOrdered" {
			eventBytes, err := json.Marshal(event)
			if err != nil {
				return errors.Wrap(err, "json.Marshal[event]")
			}

			os.baristaEventPub.Publish(ctx, eventBytes, "text/plain")
		}

		if event.Identity() == "KitchenOrdered" {
			eventBytes, err := json.Marshal(event)
			if err != nil {
				return errors.Wrap(err, "json.Marshal[event]")
			}

			os.kitchenEventPub.Publish(ctx, eventBytes, "text/plain")
		}
	}

	return nil
}

func (os *orderService) createOrderFrom(
	ctx context.Context,
	request *domain.PlaceOrderModel,
) (*domain.Order, error) {
	order := domain.NewOrder(request.OrderSource, request.LoyaltyMemberID, shared.StatusInProcess, request.Location)

	numberOfBaristaItems := len(request.BaristaItems) > 0
	numberOfKitchenItems := len(request.KitchenItems) > 0

	if numberOfBaristaItems {
		itemTypesRes, err := os.productDomainSvc.GetItemsByType(ctx, request, true)
		if err != nil {
			return nil, err
		}

		lo.ForEach(request.BaristaItems, func(item *domain.OrderItemModel, _ int) {
			find, ok := lo.Find(itemTypesRes, func(i *domain.ItemModel) bool {
				return i.ItemType == item.ItemType
			})

			if ok {
				lineItem := domain.NewLineItem(item.ItemType, item.ItemType.String(), float32(find.Price), shared.StatusInProcess, true)

				event := events.BaristaOrdered{
					OrderID:    order.ID,
					ItemLineID: lineItem.ID,
					ItemType:   item.ItemType,
				}

				order.ApplyDomain(event)

				order.LineItems = append(order.LineItems, lineItem)
			}
		})

		if err != nil {
			return nil, err
		}
	}

	if numberOfKitchenItems {
		itemTypesRes, err := os.productDomainSvc.GetItemsByType(ctx, request, false)
		if err != nil {
			return nil, err
		}

		lo.ForEach(request.KitchenItems, func(item *domain.OrderItemModel, index int) {
			find, ok := lo.Find(itemTypesRes, func(i *domain.ItemModel) bool {
				return i.ItemType == item.ItemType
			})

			if ok {
				lineItem := domain.NewLineItem(item.ItemType, item.ItemType.String(), float32(find.Price), shared.StatusInProcess, false)

				event := events.KitchenOrdered{
					OrderID:    order.ID,
					ItemLineID: lineItem.ID,
					ItemType:   item.ItemType,
				}

				order.ApplyDomain(event)

				order.LineItems = append(order.LineItems, lineItem)
			}
		})

		if err != nil {
			return nil, err
		}
	}

	return order, nil
}
