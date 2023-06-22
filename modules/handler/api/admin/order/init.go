package order

import (
	oc "github.com/berrylradianh/ecowave-go/modules/usecase/admin/order"
)

type OrderHandlerAdmin struct {
	orderUseCase oc.OrderUseCase
}

func New(orderUseCase oc.OrderUseCase) *OrderHandlerAdmin {
	return &OrderHandlerAdmin{
		orderUseCase,
	}
}
