package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(product *pe.Product) error
	CreateProductDescription(productDescription *pe.Product_Description) error
	CreateProductImage(productImage *pe.Product_Image) error
	GetAllProduct(products *[]pe.Product) ([]pe.Product, error)
	GetProductByID(productId string, product *pe.Product) (pe.Product, error)
	GetProductImageURLById(productId string, productImage *pe.Product_Image) ([]pe.Product_Image, error)
	UpdateProduct(productId string, productRequest *pe.ProductRequest) error
	UpdateProductDescription(productDescriptionID string, description string) error
	UpdateProductImage(productID string, productImage *pe.Product_Image) error
	DeleteProduct(productId string, product *pe.Product) error
	DeleteProductDescription(productDescriptionID string, productDescription *pe.Product_Description) error
	DeleteProductImage(productID string, productImages *[]pe.Product_Image) error
	SearchProductByID(productID string, product *pe.Product) (pe.Product, error)
	SearchProductByName(name string, product *[]pe.Product) ([]pe.Product, error)
	SearchProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error)
	FilterProductByStatus(status string, product *[]pe.Product) ([]pe.Product, error)
	// ImportProductFromCSV(product *pe.Product) error
}

type productRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProductRepo {
	return &productRepo{
		db,
	}
}
