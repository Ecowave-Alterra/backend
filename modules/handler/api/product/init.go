package product

import (
	up "github.com/berrylradianh/ecowave-go/modules/usecase/product"
)

type ProductHandler struct {
	productUseCase up.ProductUseCase
}

func New(productUseCase up.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase,
	}
}
