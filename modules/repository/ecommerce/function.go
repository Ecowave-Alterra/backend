package ecommerce

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (er *ecommerceRepo) GetAllProduct(products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (er *ecommerceRepo) GetProductByID(productId string, product *ep.Product) (ep.Product, error) {
	if err := er.db.Preload("ProductCategory").Where("product_id = ?", productId).First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (er *ecommerceRepo) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	var productImages []ep.ProductImage
	if err := er.db.Model(&ep.ProductImage{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, err
	}
	return productImages, nil
}

func (er *ecommerceRepo) FilterProductByCategory(category string, products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (er *ecommerceRepo) FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Order("price desc").Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (er *ecommerceRepo) FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Order("price asc").Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (er *ecommerceRepo) FilterProductByAllCategoryAndPriceMax(products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Order("price desc").Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (er *ecommerceRepo) FilterProductByAllCategoryAndPriceMin(products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Order("price asc").Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}
