package infras

import (
	"context"

	"github.com/google/wire"

	orderSvc "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/service/order"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq/publisher"
)

var (
	// definition for dependency injection
	BaristaEventPublisherSet = wire.NewSet(NewBaristaEventPublisher)
	KitchenEventPublisherSet = wire.NewSet(NewKitchenEventPublisher)
)

type (
	baristaEventPublisher struct {
		pub publisher.EventPublisher
	}
	kitchenEventPublisher struct {
		pub publisher.EventPublisher
	}
)

func NewBaristaEventPublisher(pub publisher.EventPublisher) orderSvc.BaristaEventPublisher {
	return &baristaEventPublisher{
		pub: pub,
	}
}

func (p *baristaEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *baristaEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}

func NewKitchenEventPublisher(pub publisher.EventPublisher) orderSvc.KitchenEventPublisher {
	return &kitchenEventPublisher{
		pub: pub,
	}
}

func (p *kitchenEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *kitchenEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}
