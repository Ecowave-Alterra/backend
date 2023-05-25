package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	rp "github.com/berrylradianh/ecowave-go/modules/repository/product"
)

type ProductUsecase interface {
	GetAllProducts() (*[]ep.Product, error)
}

type productUsecase struct {
	productRepository rp.ProductRepository
}

func New(productRepository rp.ProductRepository) *productUsecase {
	return &productUsecase{
		productRepository,
	}
}
