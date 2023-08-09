package usecase

import (
	"github.com/josimarz/fc-goexpert-challenge-03/internal/entity"
	"github.com/josimarz/fc-goexpert-challenge-03/pkg/events"
)

type CreateOrderInputDTO struct {
	Id    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutputDTO struct {
	Id         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	repository entity.OrderRepository
	event      events.Event
	dispatcher *events.EventDispatcher
}

func NewCreateOrderUseCase(repository entity.OrderRepository, event events.Event, dispatcher *events.EventDispatcher) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		repository,
		event,
		dispatcher,
	}
}

func (uc *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (CreateOrderOutputDTO, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)
	if err != nil {
		return CreateOrderOutputDTO{}, err
	}
	if err := order.CalculateFinalPrice(); err != nil {
		return CreateOrderOutputDTO{}, err
	}
	if err := uc.repository.Save(order); err != nil {
		return CreateOrderOutputDTO{}, err
	}
	output := CreateOrderOutputDTO{
		Id:         order.Id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	uc.event.SetPayload(output)
	uc.dispatcher.Dispatch(uc.event)
	return output, nil
}
