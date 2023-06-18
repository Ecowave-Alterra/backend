package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	er "github.com/berrylradianh/ecowave-go/modules/repository/ecommerce"
)

type EcommerceUsecase interface {
	GetProductEcommerce(products *[]ep.Product, offset, pageSize int) (*[]ee.ProductResponse, int64, error)
}

type ecommerceUseCase struct {
	ecommerceRepo er.EcommerceRepo
}

func New(ecommerceRepo er.EcommerceRepo) *ecommerceUseCase {
	return &ecommerceUseCase{
		ecommerceRepo,
	}
}
