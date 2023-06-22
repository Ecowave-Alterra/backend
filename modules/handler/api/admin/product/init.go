package product

import (
	pc "github.com/berrylradianh/ecowave-go/modules/usecase/admin/product"
)

type ProductHandler struct {
	productUseCase pc.ProductUseCase
}

func New(productUseCase pc.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase,
	}
}
