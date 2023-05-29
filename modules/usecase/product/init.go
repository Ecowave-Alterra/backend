package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	rp "github.com/berrylradianh/ecowave-go/modules/repository/product"
)

type ProductUseCase interface {
	CreateProduct(product *ep.Product) error
	CreateProductDescription(productDescription *ep.ProductDescription) error
	CreateProductImage(productImage *ep.ProductImage) error
	GetAllProduct(products *[]ep.Product) ([]ep.Product, error)
	GetProductByID(productId string, product *ep.Product) (ep.Product, error)
	GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error)
	UpdateProduct(productId string, productRequest *ep.ProductRequest) error
	UpdateProductDescription(productDescriptionID string, description string) error
	UpdateProductImage(productID string, productImage *ep.ProductImage) error
	DeleteProduct(productId string, product *ep.Product) error
	DeleteProductDescription(productDescriptionID string, productDescription *ep.ProductDescription) error
	DeleteProductImage(productID string, productImage *[]ep.ProductImage) error
	SearchProductByID(productID string, product *ep.Product) (ep.Product, error)
	SearchProductByName(name string, product *[]ep.Product) ([]ep.Product, error)
	SearchProductByCategory(category string, product *[]ep.Product) ([]ep.Product, error)
	FilterProductByStatus(status string, product *[]ep.Product) ([]ep.Product, error)
	// ImportProductFromCSV(product *pe.Product) error
}

type productUseCase struct {
	productRepo rp.ProductRepo
}

func New(productRepo rp.ProductRepo) *productUseCase {
	return &productUseCase{
		productRepo,
	}
}
