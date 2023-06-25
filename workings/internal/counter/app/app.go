package app

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/cmd/counter/config"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/events"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/infras/grpc"
	orderSvc "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/service/order"
	pkgEvents "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/pkg/events"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/postgres"
	pkgConsumer "github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq/publisher"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/proto/gen"
)

type App struct {
	Cfg       *config.Config
	PG        postgres.DBEngine
	AMQPConn  *amqp.Connection
	Publisher pkgPublisher.EventPublisher
	Consumer  pkgConsumer.EventConsumer

	BaristaOrderPub orderSvc.BaristaEventPublisher
	KitchenOrderPub orderSvc.KitchenEventPublisher

	ProductDomainSvc  grpc.ProductDomainService
	OrderService      orderSvc.OrderService
	CounterGRPCServer gen.CounterServiceServer

	baristaHandler events.BaristaOrderUpdatedEventHandler
	kitchenHandler events.KitchenOrderUpdatedEventHandler
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	amqpConn *amqp.Connection,
	publisher pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,

	baristaOrderPub orderSvc.BaristaEventPublisher,
	kitchenOrderPub orderSvc.KitchenEventPublisher,
	productDomainSvc grpc.ProductDomainService,
	orderService orderSvc.OrderService,
	counterGRPCServer gen.CounterServiceServer,

	baristaHandler events.BaristaOrderUpdatedEventHandler,
	kitchenHandler events.KitchenOrderUpdatedEventHandler,
) *App {
	return &App{
		Cfg: cfg,

		PG:        pg,
		AMQPConn:  amqpConn,
		Publisher: publisher,
		Consumer:  consumer,

		BaristaOrderPub: baristaOrderPub,
		KitchenOrderPub: kitchenOrderPub,

		ProductDomainSvc:  productDomainSvc,
		OrderService:      orderService,
		CounterGRPCServer: counterGRPCServer,

		baristaHandler: baristaHandler,
		kitchenHandler: kitchenHandler,
	}
}

func (a *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)

		switch delivery.Type {
		case "barista-order-updated":
			var payload pkgEvents.BaristaOrderUpdated

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = a.baristaHandler.Handle(ctx, &payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to delivery.Reject", err)
				}

				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		case "kitchen-order-updated":
			var payload pkgEvents.KitchenOrderUpdated

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = a.kitchenHandler.Handle(ctx, &payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to delivery.Reject", err)
				}

				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		default:
			slog.Info("default")
		}
	}

	slog.Info("Deliveries channel closed")
}
