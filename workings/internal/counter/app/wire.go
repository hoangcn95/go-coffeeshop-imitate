//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/cmd/counter/config"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/app/router"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/events/handler"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/infras"
	infrasGRPC "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/infras/grpc"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/infras/repo"
	orderSvc "github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/counter/service/order"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/postgres"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq"
	pkgConsumer "github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/hoangcn95/go-coffeeshop-imitate/workings/pkg/rabbitmq/publisher"
)

func InitApp(
	cfg *config.Config,
	dbConnStr postgres.DBConnString,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,

		infras.BaristaEventPublisherSet,
		infras.KitchenEventPublisherSet,
		infrasGRPC.ProductGRPCClientSet,
		router.CounterGRPCServerSet,
		repo.RepositorySet,
		orderSvc.OrderServiceSet,
		handler.BaristaOrderUpdatedEventHandlerSet,
		handler.KitchenOrderUpdatedEventHandlerSet,
	))
}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
