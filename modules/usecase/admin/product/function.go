package product

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pc *productUseCase) CreateProduct(product *pe.Product) error {
	for {
		productId := randomid.GenerateRandomID()

		exists, err := pc.productRepo.CheckProductExist(productId)
		if err != nil {
			return err
		}
		if !exists {
			product.ProductId = productId
			break
		}
	}
	return pc.productRepo.CreateProduct(product)
}

func (pc *productUseCase) CreateProductImage(productImage *pe.ProductImage) error {
	return pc.productRepo.CreateProductImage(productImage)
}

func (pc *productUseCase) GetAllProduct(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	product, count, err := pc.productRepo.GetAllProduct(products, offset, pageSize)
	return product, count, err
}

func (pc *productUseCase) GetAllProductNoPagination(products *[]pe.Product) ([]pe.Product, error) {
	return pc.productRepo.GetAllProductNoPagination(products)
}

func (pc *productUseCase) GetProductByID(productId string, product *pe.Product) (*pe.Product, int64, float64, error) {
	return pc.productRepo.GetProductByID(productId, product)
}

func (pc *productUseCase) UpdateProduct(productId string, productRequest *pe.ProductRequest) error {
	return pc.productRepo.UpdateProduct(productId, productRequest)
}

func (pc *productUseCase) UpdateProductStock(productId string, stock uint) error {
	return pc.productRepo.UpdateProductStock(productId, stock)
}

func (pc *productUseCase) CountProductImage(productId string) (int, error) {
	return pc.productRepo.CountProductImage(productId)
}

func (pc *productUseCase) UpdateProductImage(idImage int, productId string, productImageUrl string) error {
	return pc.productRepo.UpdateProductImage(idImage, productId, productImageUrl)
}

func (pc *productUseCase) DeleteProduct(productId string, product *pe.Product) error {
	return pc.productRepo.DeleteProduct(productId, product)
}

func (pc *productUseCase) SearchProduct(search, filter string, offset, pageSize int) (*[]pe.Product, int64, error) {
	products, count, err := pc.productRepo.SearchProduct(search, filter, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}
