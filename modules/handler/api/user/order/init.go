package order

import (
	uo "github.com/berrylradianh/ecowave-go/modules/usecase/user/order"
)

type OrderHandler struct {
	orderUsecase uo.OderUsecase
}

func New(informationUsecase uo.OderUsecase) *OrderHandler {
	return &OrderHandler{
		informationUsecase,
	}
}
