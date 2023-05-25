package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (productUsecase *productUsecase) GetAllProducts() (*[]ep.Product, error) {
	products, err := productUsecase.productRepository.GetAllProducts()

	return products, err
}
