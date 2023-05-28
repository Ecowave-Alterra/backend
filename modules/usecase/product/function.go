package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pc *productUseCase) CreateProduct(product *pe.Product) error {
	return pc.productRepo.CreateProduct(product)
}

func (pc *productUseCase) CreateProductDescription(productDescription *pe.Product_Description) error {
	return pc.productRepo.CreateProductDescription(productDescription)
}

func (pc *productUseCase) CreateProductImage(productImage *pe.Product_Image) error {
	return pc.productRepo.CreateProductImage(productImage)
}

func (pc *productUseCase) GetAllProduct(products *[]pe.Product) ([]pe.Product, error) {
	return pc.productRepo.GetAllProduct(products)
}

func (pc *productUseCase) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	return pc.productRepo.GetProductByID(productId, product)
}

func (pc *productUseCase) GetProductImageURLById(productId string, productImage *pe.Product_Image) ([]pe.Product_Image, error) {
	return pc.productRepo.GetProductImageURLById(productId, productImage)
}

func (pc *productUseCase) UpdateProduct(productId string, productRequest *pe.ProductRequest) error {
	return pc.productRepo.UpdateProduct(productId, productRequest)
}

func (pc *productUseCase) UpdateProductDescription(productDescriptionID string, description string) error {
	return pc.productRepo.UpdateProductDescription(productDescriptionID, description)
}

func (pc *productUseCase) UpdateProductImage(productID string, productImage *pe.Product_Image) error {
	return pc.productRepo.UpdateProductImage(productID, productImage)
}

func (pc *productUseCase) DeleteProduct(productId string, product *pe.Product) error {
	return pc.productRepo.DeleteProduct(productId, product)
}

func (pc *productUseCase) DeleteProductDescription(productDescriptionID string, productDescription *pe.Product_Description) error {
	return pc.productRepo.DeleteProductDescription(productDescriptionID, productDescription)
}

func (pc *productUseCase) DeleteProductImage(productID string, productImages *[]pe.Product_Image) error {
	return pc.productRepo.DeleteProductImage(productID, productImages)
}

func (pc *productUseCase) SearchProductByID(productID string, product *pe.Product) (pe.Product, error) {
	return pc.productRepo.SearchProductByID(productID, product)
}

func (pc *productUseCase) SearchProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
	return pc.productRepo.SearchProductByName(name, product)
}

func (pc *productUseCase) SearchProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
	return pc.productRepo.SearchProductByCategory(category, product)
}

func (pc *productUseCase) FilterProductByStatus(status string, product *[]pe.Product) ([]pe.Product, error) {
	return pc.productRepo.FilterProductByStatus(status, product)
}

// func (pc *productUseCase) ImportProductFromCSV(product *pe.Product) error {
// 	return pc.productRepo.ImportProductFromCSV(product)
// }
