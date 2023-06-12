package order

import (
	ro "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
)

type OderUsecase interface {
	GetOrder(id string, idUser uint, offset int, pageSize int) (interface{}, int64, error)
	OrderDetail(id uint) (interface{}, error)
	ConfirmOrder(id uint) error
	CancelOrder(id uint, canceledReason string) error
}

type orderUsecase struct {
	orderRepo ro.OrderRepo
}

func New(orderRepo ro.OrderRepo) *orderUsecase {
	return &orderUsecase{
		orderRepo,
	}
}
