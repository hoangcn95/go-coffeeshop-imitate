package domain

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/events"
	shared "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/shared_kernel"
)

type Order struct {
	shared.AggregateRoot
	ID              uuid.UUID
	OrderSource     shared.OrderSource
	LoyaltyMemberID uuid.UUID
	OrderStatus     shared.Status
	Location        shared.Location
	LineItems       []*LineItem
}

func NewOrder(
	orderSource shared.OrderSource,
	loyaltyMemberID uuid.UUID,
	orderStatus shared.Status,
	location shared.Location,
) *Order {
	return &Order{
		ID:              uuid.New(),
		OrderSource:     orderSource,
		LoyaltyMemberID: loyaltyMemberID,
		OrderStatus:     orderStatus,
		Location:        location,
	}
}

func (o *Order) Apply(event *events.OrderUp) error {
	if len(o.LineItems) == 0 {
		return nil // we dont do anything
	}

	_, index, ok := lo.FindIndexOf(o.LineItems, func(i *LineItem) bool {
		return i.ItemType == event.ItemType
	})

	if !ok {
		return ErrItemNotFound
	}

	o.LineItems[index].ItemStatus = shared.StatusFulfilled

	if checkFulfilledStatus(o.LineItems) {
		o.OrderStatus = shared.StatusFulfilled
	}

	return nil
}

func checkFulfilledStatus(lineItems []*LineItem) bool {
	for _, item := range lineItems {
		if item.ItemStatus != shared.StatusFulfilled {
			return false
		}
	}

	return true
}
