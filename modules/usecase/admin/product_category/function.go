package product_category

import (
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pcc *productCategoryUsecase) CreateProductCategory(productCategory *pe.ProductCategory) (bool, error) {
	if err := vld.Validation(productCategory); err != nil {
		return false, err
	}

	available, _ := pcc.productCategoryRepo.IsProductCategoryAvailable(productCategory, productCategory.Name)
	return available, pcc.productCategoryRepo.CreateProductCategory(productCategory)
}

func (pcc *productCategoryUsecase) UpdateProductCategory(productCategory *pe.ProductCategory, id int) (bool, error) {
	available, _ := pcc.productCategoryRepo.IsProductCategoryAvailable(productCategory, productCategory.Name)
	return available, pcc.productCategoryRepo.UpdateProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) DeleteProductCategory(productCategory *pe.ProductCategory, id int) error {
	return pcc.productCategoryRepo.DeleteProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) GetAllProductCategory(offset, pageSize int) (*[]pe.ProductCategory, int64, error) {
	productCategories, count, err := pcc.productCategoryRepo.GetAllProductCategory(offset, pageSize)
	return productCategories, count, err
}

func (pcc *productCategoryUsecase) SearchingProductCategoryByName(productCategory *[]pe.ProductCategory, name string) (bool, error) {
	return pcc.productCategoryRepo.SearchingProductCategoryByName(productCategory, name)
}
