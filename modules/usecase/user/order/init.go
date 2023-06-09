package order

import (
	ro "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
)

type OderUsecase interface {
	GetOrder(id string, idUser uint) (interface{}, error)
	OrderDetail(id uint) (interface{}, error)
}

type orderUsecase struct {
	orderRepo ro.OrderRepo
}

func New(orderRepo ro.OrderRepo) *orderUsecase {
	return &orderUsecase{
		orderRepo,
	}
}
