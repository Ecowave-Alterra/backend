package product_category

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pcr *productCategoryRepo) CreateProductCategory(productCategory *pe.ProductCategory) error {
	if err := pcr.db.Save(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (pcr *productCategoryRepo) UpdateProductCategory(productCategory *pe.ProductCategory, id int) error {
	if err := pcr.db.Where("id = ?", id).Updates(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (pcr *productCategoryRepo) DeleteProductCategory(productCategory *pe.ProductCategory, id int) error {
	if err := pcr.db.Where("id = ?", id).Delete(&productCategory).Error; err != nil {
		return err
	}

	return nil
}
func (pcr *productCategoryRepo) GetAllProductCategory(offset, pageSize int) (*[]pe.ProductCategory, int64, error) {
	var productCategories []pe.ProductCategory
	var count int64
	if err := pcr.db.Model(&pe.ProductCategory{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := pcr.db.Offset(offset).Limit(pageSize).Find(&productCategories).Error; err != nil {
		return nil, 0, err
	}

	return &productCategories, count, nil
}

func (pcr *productCategoryRepo) SearchingProductCategoryByName(productCategory *[]pe.ProductCategory, name string) (bool, error) {
	result := pcr.db.Where("name LIKE ?", "%"+name+"%").Find(&productCategory)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (pcr *productCategoryRepo) IsProductCategoryAvailable(productCategory *pe.ProductCategory, name string) (bool, error) {
	result := pcr.db.Where("name = ?", name).Find(&productCategory)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
