package handlers

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/events"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/service/order"
	pkgEvents "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/events"
)

type baristaOrderUpdatedEventHandler struct {
	orderRepo order.OrderRepo
}

var _ events.BaristaOrderUpdatedEventHandler = (*baristaOrderUpdatedEventHandler)(nil)

// definitions for dependency injection
var BaristaOrderUpdatedEventHandlerSet = wire.NewSet(NewBaristaOrderUpdatedEventHandler)

func NewBaristaOrderUpdatedEventHandler(orderRepo order.OrderRepo) events.BaristaOrderUpdatedEventHandler {
	return &baristaOrderUpdatedEventHandler{
		orderRepo: orderRepo,
	}
}

func (h *baristaOrderUpdatedEventHandler) Handle(ctx context.Context, e *pkgEvents.BaristaOrderUpdated) error {
	order, err := h.orderRepo.GetByID(ctx, e.OrderID)
	if err != nil {
		return errors.Wrap(err, "orderRepo.GetByID")
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
