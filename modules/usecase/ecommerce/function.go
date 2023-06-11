package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (ec *ecommerceUseCase) GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	products, count, err := ec.ecommerceRepo.GetAllProduct(products, offset, pageSize)
	return products, count, err
}

func (ec *ecommerceUseCase) GetProductByID(productId string) ([]ee.QueryResponse, error) {
	rq, err := ec.ecommerceRepo.GetProductByID(productId)
	return rq, err
}

func (ec *ecommerceUseCase) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	return ec.ecommerceRepo.GetProductImageURLById(productId, productImage)
}

func (ec *ecommerceUseCase) FilterProductByCategory(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	products, count, err := ec.ecommerceRepo.FilterProductByCategory(category, products, offset, pageSize)
	return products, count, err
}

func (ec *ecommerceUseCase) FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	products, count, err := ec.ecommerceRepo.FilterProductByCategoryAndPriceMax(category, products, offset, pageSize)
	return products, count, err
}

func (ec *ecommerceUseCase) FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	products, count, err := ec.ecommerceRepo.FilterProductByCategoryAndPriceMin(category, products, offset, pageSize)
	return products, count, err
}

func (ec *ecommerceUseCase) FilterProductByAllCategoryAndPriceMax(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	products, count, err := ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMax(products, offset, pageSize)
	return products, count, err
}

func (ec *ecommerceUseCase) FilterProductByAllCategoryAndPriceMin(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	products, count, err := ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMin(products, offset, pageSize)
	return products, count, err
}
