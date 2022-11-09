package usecase

import "github.com/pr02nl/gointensivo/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (uc *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := uc.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDTO{Total: total}, nil
}
