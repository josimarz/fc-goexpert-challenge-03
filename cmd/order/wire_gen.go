// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/entity"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/event"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/infra/database"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/infra/web"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/usecase"
	"github.com/josimarz/fc-goexpert-challenge-03/pkg/events"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, dispatcher *events.EventDispatcher) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreatedEvent := event.NewOrderCreatedEvent()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreatedEvent, dispatcher)
	return createOrderUseCase
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)
	return listOrdersUseCase
}

func NewWebOrderHandler(db *sql.DB, dispatcher *events.EventDispatcher) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreatedEvent := event.NewOrderCreatedEvent()
	webOrderHandler := web.NewWebOrderHandler(orderRepository, orderCreatedEvent, dispatcher)
	return webOrderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepository), new(*database.OrderRepository)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreatedEvent, wire.Bind(new(events.Event), new(*event.OrderCreatedEvent)))
