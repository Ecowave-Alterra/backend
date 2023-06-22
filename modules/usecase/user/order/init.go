package order

import (
	eo "github.com/berrylradianh/ecowave-go/modules/entity/order"
	ro "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
)

type OderUsecase interface {
	GetOrder(filter string, idUser uint, offset int, pageSize int) (interface{}, int64, error)
	Tracking(resi string, courier string) (interface{}, error)
	ConfirmOrder(eo.ConfirmOrder) error
	CancelOrder(co eo.CanceledOrder) error
}

type orderUsecase struct {
	orderRepo ro.OrderRepo
}

func New(orderRepo ro.OrderRepo) *orderUsecase {
	return &orderUsecase{
		orderRepo,
	}
}
