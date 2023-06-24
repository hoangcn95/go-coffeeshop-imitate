package domain

import (
	"github.com/google/uuid"

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
