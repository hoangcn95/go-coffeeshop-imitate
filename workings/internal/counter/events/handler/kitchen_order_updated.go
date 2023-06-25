package handler

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/events"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/service/order"
	pkgEvents "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/events"
)

type kitchenOrderUpdatedEventHandler struct {
	orderRepo order.OrderRepo
}

var _ events.KitchenOrderUpdatedEventHandler = (*kitchenOrderUpdatedEventHandler)(nil)

// definition dependency injection
var KitchenOrderUpdatedEventHandlerSet = wire.NewSet(NewKitchenOrderUpdatedEventHandler)

func NewKitchenOrderUpdatedEventHandler(orderRepo order.OrderRepo) events.KitchenOrderUpdatedEventHandler {
	return &kitchenOrderUpdatedEventHandler{
		orderRepo: orderRepo,
	}
}

func (h *kitchenOrderUpdatedEventHandler) Handle(ctx context.Context, e *pkgEvents.KitchenOrderUpdated) error {
	order, err := h.orderRepo.GetByID(ctx, e.OrderID)
	if err != nil {
		return errors.Wrap(err, "orderRepo.GetOrderByID")
	}

	orderUp := pkgEvents.OrderUp{
		OrderID:    e.OrderID,
		ItemLineID: e.ItemLineID,
		Name:       e.Name,
		ItemType:   e.ItemType,
		TimeUp:     e.TimeUp,
		MadeBy:     e.MadeBy,
	}

	if err = order.Apply(&orderUp); err != nil {
		return errors.Wrap(err, "order.Apply")
	}

	_, err = h.orderRepo.Update(ctx, order)
	if err != nil {
		return errors.Wrap(err, "orderRepo.Update")
	}

	return nil
}
