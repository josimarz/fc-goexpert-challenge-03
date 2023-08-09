//go:build wireinject
// +build wireinject

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

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepository), new(*database.OrderRepository)),
)
var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreatedEvent,
	wire.Bind(new(events.Event), new(*event.OrderCreatedEvent)),
)

func NewCreateOrderUseCase(db *sql.DB, dispatcher *events.EventDispatcher) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, dispatcher *events.EventDispatcher) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
