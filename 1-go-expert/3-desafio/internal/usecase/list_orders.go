package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersInputDTO struct {
}

type OrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutputDTO struct {
	Orders []OrderOutput `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute(input ListOrdersInputDTO) (ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.All()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var orderOutputs []OrderOutput
	for _, order := range orders {
		orderOutput := OrderOutput{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		orderOutputs = append(orderOutputs, orderOutput)
	}

	return ListOrdersOutputDTO{
		Orders: orderOutputs,
	}, nil
}
