package product_category

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
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

func (pcr *productCategoryRepo) GetAllProductCategoryNoPagination() (*[]pe.ProductCategory, error) {
	var productCategories []pe.ProductCategory
	if err := pcr.db.Preload("Products ").Find(&productCategories).Error; err != nil {
		return nil, err
	}

	return &productCategories, nil
}

func (pcr *productCategoryRepo) GetProductCategoryById(id int) (*pe.ProductCategory, error) {
	var productCategory pe.ProductCategory
	if err := pcr.db.Where("id = ?", id).Preload("Products").First(&productCategory).Error; err != nil {
		return nil, err
	}

	return &productCategory, nil
}

func (pcr *productCategoryRepo) GetAllProductCategory(offset, pageSize int) (*[]pe.ProductCategory, int64, error) {
	var productCategories []pe.ProductCategory
	var count int64
	if err := pcr.db.Model(&pe.ProductCategory{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := pcr.db.Preload("Products").Offset(offset).Limit(pageSize).Find(&productCategories).Error; err != nil {
		return nil, 0, err
	}

	return &productCategories, count, nil
}

func (pcr *productCategoryRepo) SearchingProductCategoryByName(name string, offset, pageSize int) (*[]pe.ProductCategory, int64, error) {
	var productCategories []pe.ProductCategory
	var count int64

	if err := pcr.db.Model(&pe.ProductCategory{}).Where("category LIKE ?", "%"+name+"%").
		Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := pcr.db.Model(&pe.ProductCategory{}).Where("category LIKE ?", "%"+name+"%").
		Count(&count).
		Offset(offset).
		Limit(pageSize).
		Preload("Products").
		Find(&productCategories).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return &productCategories, count, nil
}

func (pcr *productCategoryRepo) IsProductCategoryAvailable(productCategory *pe.ProductCategory, name string) (bool, error) {
	result := pcr.db.Where("category = ?", name).Find(&productCategory)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
