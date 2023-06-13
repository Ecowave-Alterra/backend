package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	rp "github.com/berrylradianh/ecowave-go/modules/repository/admin/product"
)

type ProductUseCase interface {
	CreateProduct(product *pe.Product) error
	CreateProductImage(productImage *pe.ProductImage) error
	GetAllProduct(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error)
	GetAllProductNoPagination(products *[]pe.Product) ([]pe.Product, error)
	GetProductByID(productId string, product *pe.Product) (pe.Product, error)
	GetProductImageURLById(productId string, productImage *pe.ProductImage) ([]pe.ProductImage, error)
	UpdateProduct(productId string, productRequest *pe.ProductRequest) error
	UpdateProductStock(productId string, stock uint) error
	DeleteProduct(productId string, product *pe.Product) error
	DeleteProductImage(productID string, productImage *[]pe.ProductImage) error
	DeleteProductImageByID(ProductImageID string, productImage *pe.ProductImage) error
	SearchProductByID(productID string, product *pe.Product) (pe.Product, error)
	SearchProductByName(name string, product *[]pe.Product) ([]pe.Product, error)
	SearchProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error)
	FilterProductByStatus(status string, product *[]pe.Product) ([]pe.Product, error)
}

type productUseCase struct {
	productRepo rp.ProductRepo
}

func New(productRepo rp.ProductRepo) *productUseCase {
	return &productUseCase{
		productRepo,
	}
}
