package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	rp "github.com/berrylradianh/ecowave-go/modules/repository/admin/product"
)

type ProductUseCase interface {
	CreateProduct(product *pe.Product) error
	GetProductByID(productId string, product *pe.Product) (pe.Product, error)
	SearchProduct(search, filter string, offset, pageSize int) (*[]pe.Product, int64, error)
	GetAllProduct(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error)
	CreateProductImage(productImage *pe.ProductImage) error
	UpdateProduct(productId string, productRequest *pe.ProductRequest) error
	UpdateProductStock(productId string, stock uint) error
	DeleteProductImage(productID string, productImage *[]pe.ProductImage) error
	DeleteProductImageByID(ProductImageID uint, productImage *pe.ProductImage) error
	DeleteProduct(productId string, product *pe.Product) error
	GetAllProductNoPagination(products *[]pe.Product) ([]pe.Product, error)
	GetProductImageURLById(productId string, productImage *pe.ProductImage) ([]pe.ProductImage, error)
}

type productUseCase struct {
	productRepo rp.ProductRepo
}

func New(productRepo rp.ProductRepo) *productUseCase {
	return &productUseCase{
		productRepo,
	}
}
