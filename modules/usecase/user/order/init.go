package order

import (
	ro "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
)

type OderUsecase interface {
	GetOrder(filter string, idUser uint, offset int, pageSize int) (interface{}, int64, error)
	// OrderDetail(id uint) (interface{}, error)
	Tracking(resi string, courier string) (interface{}, error)
	ConfirmOrder(id string) error
	CancelOrder(id string, canceledReason string) error
}

type orderUsecase struct {
	orderRepo ro.OrderRepo
}

func New(orderRepo ro.OrderRepo) *orderUsecase {
	return &orderUsecase{
		orderRepo,
	}
}