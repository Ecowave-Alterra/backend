package ecommerce

import (
	ec "github.com/berrylradianh/ecowave-go/modules/usecase/user/ecommerce"
)

type EcommerceHandler struct {
	ecommerceUseCase ec.EcommerceUsecase
}

func New(ecommerceUseCase ec.EcommerceUsecase) *EcommerceHandler {
	return &EcommerceHandler{
		ecommerceUseCase,
	}
}
