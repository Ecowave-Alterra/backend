package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(product *ep.Product) error
	CreateProductImage(productImage *ep.ProductImage) error
	GetAllProduct(products *[]ep.Product) ([]ep.Product, error)
	GetProductByID(productId string, product *ep.Product) (ep.Product, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	UpdateProduct(productId string, productRequest *ep.ProductRequest) error
	UpdateProductImage(productID string, productImage *ep.ProductImage) error
	DeleteProduct(productId string, product *ep.Product) error
	DeleteProductImage(productID string, productImages *[]ep.ProductImage) error
	SearchProductByID(productID string, product *ep.Product) (ep.Product, error)
	SearchProductByName(name string, product *[]ep.Product) ([]ep.Product, error)
	SearchProductByCategory(category string, product *[]ep.Product) ([]ep.Product, error)
	FilterProductByStatus(status string, product *[]ep.Product) ([]ep.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProductRepo {
	return &productRepo{
		db,
	}
}
