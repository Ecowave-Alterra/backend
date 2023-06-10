package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (ec *ecommerceUseCase) GetAllProduct(products *[]ep.Product) ([]ep.Product, error) {
	return ec.ecommerceRepo.GetAllProduct(products)
}

func (ec *ecommerceUseCase) GetProductByID(productId string) ([]ee.QueryResponse, error) {
	rq, err := ec.ecommerceRepo.GetProductByID(productId)
	return rq, err
}

func (ec *ecommerceUseCase) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	return ec.ecommerceRepo.GetProductImageURLById(productId, productImage)
}

func (ec *ecommerceUseCase) FilterProductByCategory(category string, products *[]ep.Product) ([]ep.Product, error) {
	return ec.ecommerceRepo.FilterProductByCategory(category, products)
}

func (ec *ecommerceUseCase) FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product) ([]ep.Product, error) {
	return ec.ecommerceRepo.FilterProductByCategoryAndPriceMax(category, products)
}

func (ec *ecommerceUseCase) FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product) ([]ep.Product, error) {
	return ec.ecommerceRepo.FilterProductByCategoryAndPriceMin(category, products)
}

func (ec *ecommerceUseCase) FilterProductByAllCategoryAndPriceMax(products *[]ep.Product) ([]ep.Product, error) {
	return ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMax(products)
}

func (ec *ecommerceUseCase) FilterProductByAllCategoryAndPriceMin(products *[]ep.Product) ([]ep.Product, error) {
	return ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMin(products)
}
