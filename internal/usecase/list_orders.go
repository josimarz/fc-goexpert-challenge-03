package usecase

import "github.com/josimarz/fc-goexpert-challenge-03/internal/entity"

type ListOrdersOutputDTO []entity.Order

type ListOrdersUseCase struct {
	repository entity.OrderRepository
}

func NewListOrdersUseCase(repository entity.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{repository}
}

func (uc *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := uc.repository.FindAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}
	return orders, nil
}
