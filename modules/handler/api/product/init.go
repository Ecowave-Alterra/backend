package product

import (
	up "github.com/berrylradianh/ecowave-go/modules/usecase/product"
)

type ProductHandler struct {
	productUsecase up.ProductUsecase
}

func New(productUsecase up.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase,
	}
}
