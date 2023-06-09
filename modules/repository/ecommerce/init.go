package ecommerce

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"gorm.io/gorm"
)

type EcommerceRepo interface {
	GetAllProduct(products *[]ep.Product) ([]ep.Product, error)
	GetProductByID(productId string, product *ep.Product) (ep.Product, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	FilterProductByCategory(category string, products *[]ep.Product) ([]ep.Product, error)
	FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product) ([]ep.Product, error)
	FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product) ([]ep.Product, error)
	FilterProductByAllCategoryAndPriceMax(products *[]ep.Product) ([]ep.Product, error)
	FilterProductByAllCategoryAndPriceMin(products *[]ep.Product) ([]ep.Product, error)
}

type ecommerceRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) EcommerceRepo {
	return &ecommerceRepo{
		db,
	}
}
