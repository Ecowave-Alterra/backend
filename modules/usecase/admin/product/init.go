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
	GetProductByID(productId string, product *pe.Product) (*pe.Product, int64, float64, error)
	UpdateProduct(productId string, productRequest *pe.ProductRequest) error
	UpdateProductStock(productId string, stock uint) error
	CountProductImage(productId string) (int, error)
	UpdateProductImage(idImage int, productId string, productImageUrl string) error
	DeleteProduct(productId string, product *pe.Product) error
	SearchProduct(search, filter string, offset, pageSize int) (*[]pe.Product, int64, error)
}

type productUseCase struct {
	productRepo rp.ProductRepo
}

func New(productRepo rp.ProductRepo) *productUseCase {
	return &productUseCase{
		productRepo,
	}
}
