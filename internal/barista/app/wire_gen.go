// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/thangchung/go-coffeeshop/cmd/barista/config"
	"github.com/thangchung/go-coffeeshop/internal/barista/eventhandlers"
	"github.com/thangchung/go-coffeeshop/pkg/postgres"
	"github.com/thangchung/go-coffeeshop/pkg/rabbitmq"
	"github.com/thangchung/go-coffeeshop/pkg/rabbitmq/consumer"
	"github.com/thangchung/go-coffeeshop/pkg/rabbitmq/publisher"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, dbConnStr postgres.DBConnString, rabbitMQConnStr rabbitmq.RabbitMQConnStr) (*App, func(), error) {
	dbEngine, cleanup, err := dbEngineFunc(dbConnStr)
	if err != nil {
		return nil, nil, err
	}
	connection, cleanup2, err := rabbitMQFunc(rabbitMQConnStr)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	eventPublisher, err := publisher.NewPublisher(connection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	eventConsumer, err := consumer.NewConsumer(connection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	baristaOrderedEventHandler := eventhandlers.NewBaristaOrderedEventHandler(dbEngine, eventPublisher)
	app := New(cfg, dbEngine, connection, eventPublisher, eventConsumer, baristaOrderedEventHandler)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp091.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
