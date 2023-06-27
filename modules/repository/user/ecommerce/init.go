package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"gorm.io/gorm"
)

type EcommerceRepo interface {
	GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error)
	GetProductByID(productId string) (bool, []ee.ReviewResponse, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	AvgRating(productId string) (float64, error)
}

type ecommerceRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) EcommerceRepo {
	return &ecommerceRepo{
		db,
	}
}
