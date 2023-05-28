package productcategory

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pcr *productCategoryRepo) CreateProductCategory(productCategory *pct.ProductCategory) error {
	if err := pcr.db.Save(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (pcr *productCategoryRepo) UpdateProductCategory(productCategory *pct.ProductCategory, id int) error {
	if err := pcr.db.Where("id = ?", id).Updates(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (pcr *productCategoryRepo) DeleteProductCategory(productCategory *pct.ProductCategory, id int) error {
	if err := pcr.db.Where("id = ?", id).Delete(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (pcr *productCategoryRepo) GetAllProductCategory(productCategory *[]pct.ProductCategory) error {
	if err := pcr.db.Find(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (pcr *productCategoryRepo) SearchingProductCategoryByName(productCategory *[]pct.ProductCategory, name string) (bool, error) {
	result := pcr.db.Where("name LIKE ?", "%"+name+"%").Find(&productCategory)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (pcr *productCategoryRepo) IsProductCategoryAvailable(productCategory *pct.ProductCategory, name string) (bool, error) {
	result := pcr.db.Where("name = ?", name).Find(&productCategory)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
