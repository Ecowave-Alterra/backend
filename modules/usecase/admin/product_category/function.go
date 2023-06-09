package product_category

import pct "github.com/berrylradianh/ecowave-go/modules/entity/product"

func (pcc *productCategoryUsecase) CreateProductCategory(productCategory *pct.ProductCategory) (bool, error) {
	available, _ := pcc.productCategoryRepo.IsProductCategoryAvailable(productCategory, productCategory.Name)
	return available, pcc.productCategoryRepo.CreateProductCategory(productCategory)
}

func (pcc *productCategoryUsecase) UpdateProductCategory(productCategory *pct.ProductCategory, id int) (bool, error) {
	available, _ := pcc.productCategoryRepo.IsProductCategoryAvailable(productCategory, productCategory.Name)
	return available, pcc.productCategoryRepo.UpdateProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) DeleteProductCategory(productCategory *pct.ProductCategory, id int) error {
	return pcc.productCategoryRepo.DeleteProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) GetAllProductCategory(offset, pageSize int) (*[]pct.ProductCategory, int64, error) {
	productCategories, count, err := pcc.productCategoryRepo.GetAllProductCategory(offset, pageSize)
	return productCategories, count, err
}

func (pcc *productCategoryUsecase) SearchingProductCategoryByName(productCategory *[]pct.ProductCategory, name string) (bool, error) {
	return pcc.productCategoryRepo.SearchingProductCategoryByName(productCategory, name)
}
