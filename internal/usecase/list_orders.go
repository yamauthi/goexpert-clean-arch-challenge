package usecase

import (
	"github.com/yamauthi/goexpert-clean-arch-challenge/internal/entity"
)

type OrdersListOutputDTO struct {
	Orders []OrderOutputDTO
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (l *ListOrdersUseCase) Execute() (OrdersListOutputDTO, error) {
	orders, err := l.OrderRepository.List()

	if err != nil {
		return OrdersListOutputDTO{}, err
	}

	var ordersOutput []OrderOutputDTO

	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return OrdersListOutputDTO{Orders: ordersOutput}, nil
}
