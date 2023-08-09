package service

import (
	"context"

	"github.com/josimarz/fc-goexpert-challenge-03/internal/infra/grpc/pb"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	createOrderUseCase *usecase.CreateOrderUseCase
	listOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		createOrderUseCase: createOrderUseCase,
		listOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	input := usecase.CreateOrderInputDTO{
		Id:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.createOrderUseCase.Execute(input)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.Id,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Empty) (*pb.OrderList, error) {
	output, err := s.listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var items []*pb.Order
	for _, order := range output {
		item := &pb.Order{
			Id:         order.Id,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		items = append(items, item)
	}
	return &pb.OrderList{Orders: items}, nil
}
