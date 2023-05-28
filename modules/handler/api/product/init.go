package product

import pc "github.com/berrylradianh/ecowave-go/modules/usecase/product"

type ProductHandler struct {
	productUC pc.ProductUseCase
}

func New(productUC pc.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUC,
	}
}
