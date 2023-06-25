package events

import (
	"context"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/events"
)

type (
	BaristaOrderUpdatedEventHandler interface {
		Handle(context.Context, *events.BaristaOrderUpdated) error
	}

	KitchenOrderUpdatedEventHandler interface {
		Handle(context.Context, *events.KitchenOrderUpdated) error
	}
)
