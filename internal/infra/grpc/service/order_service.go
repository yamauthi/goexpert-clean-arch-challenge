package service

import (
	"context"

	"github.com/yamauthi/goexpert-clean-arch-challenge/internal/infra/grpc/pb"
	"github.com/yamauthi/goexpert-clean-arch-challenge/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:         output.ID,
			Price:      float32(output.Price),
			Tax:        float32(output.Tax),
			FinalPrice: float32(output.FinalPrice),
		},
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, b *pb.Blank) (*pb.ListOrdersResponse, error) {
	list, err := s.ListOrdersUseCase.Execute()

	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.Order
	for _, o := range list.Orders {
		order := &pb.Order{
			Id:         o.ID,
			Price:      float32(o.Price),
			Tax:        float32(o.Tax),
			FinalPrice: float32(o.FinalPrice),
		}

		ordersResponse = append(ordersResponse, order)
	}

	return &pb.ListOrdersResponse{Orders: ordersResponse}, nil
}
