package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	er "github.com/berrylradianh/ecowave-go/modules/repository/ecommerce"
)

type EcommerceUsecase interface {
	GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	GetProductByID(productId string) ([]ee.QueryResponse, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	FilterProductByCategory(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	FilterProductByAllCategoryAndPriceMax(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	FilterProductByAllCategoryAndPriceMin(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
}

type ecommerceUseCase struct {
	ecommerceRepo er.EcommerceRepo
}

func New(ecommerceRepo er.EcommerceRepo) *ecommerceUseCase {
	return &ecommerceUseCase{
		ecommerceRepo,
	}
}
