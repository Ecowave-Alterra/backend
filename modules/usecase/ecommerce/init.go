package ecommerce

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	er "github.com/berrylradianh/ecowave-go/modules/repository/ecommerce"
)

type EcommerceUsecase interface {
	GetAllProduct(products *[]ep.Product) ([]ep.Product, error)
	GetProductByID(productId string, product *ep.Product) (ep.Product, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	FilterProductByCategory(category string, products *[]ep.Product) ([]ep.Product, error)
	FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product) ([]ep.Product, error)
	FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product) ([]ep.Product, error)
	FilterProductByAllCategoryAndPriceMax(products *[]ep.Product) ([]ep.Product, error)
	FilterProductByAllCategoryAndPriceMin(products *[]ep.Product) ([]ep.Product, error)
}

type ecommerceUseCase struct {
	ecommerceRepo er.EcommerceRepo
}

func New(ecommerceRepo er.EcommerceRepo) *ecommerceUseCase {
	return &ecommerceUseCase{
		ecommerceRepo,
	}
}
