package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (r *productRepo) CreateProduct(product *pe.Product) error {
	if err := r.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) CreateProductDescription(productDescription *pe.Product_Description) error {
	if err := r.db.Save(&productDescription).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) CreateProductImage(productImage *pe.Product_Image) error {
	if err := r.db.Save(&productImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) GetAllProduct(products *[]pe.Product) ([]pe.Product, error) {
	if err := r.db.
		Preload("Product_Category").Preload("Product_Description").
		Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (r *productRepo) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	if err := r.db.
		Preload("Product_Category").
		Preload("Product_Description").
		First(&product, productId).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (r *productRepo) GetProductImageURLById(productId string, productImage *pe.Product_Image) ([]pe.Product_Image, error) {
	var productImages []pe.Product_Image
	if err := r.db.Model(&pe.Product_Image{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, err
	}
	return productImages, nil
}

func (r *productRepo) UpdateProduct(productId string, req *pe.ProductRequest) error {
	if err := r.db.Model(&pe.Product{}).Where("id = ?", productId).Updates(pe.Product{Product_category_id: req.Product_category_id, Name: req.Name, Stock: req.Stock, Price: req.Price, Status: req.Status}).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) UpdateProductDescription(productDescriptionID string, description string) error {
	if err := r.db.Model(&pe.Product_Description{}).Where("id = ?", productDescriptionID).Updates(pe.Product_Description{Description: description}).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) UpdateProductImage(productID string, productImage *pe.Product_Image) error {
	return nil
}

func (r *productRepo) DeleteProduct(productId string, product *pe.Product) error {
	if err := r.db.Delete(&product, productId).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) DeleteProductDescription(productDescriptionID string, productDescription *pe.Product_Description) error {
	if err := r.db.Delete(&productDescription, productDescriptionID).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) DeleteProductImage(productID string, productImages *[]pe.Product_Image) error {
	if err := r.db.Where("product_id = ?", productID).Delete(&productImages).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) SearchProductByID(productID string, product *pe.Product) (pe.Product, error) {
	if err := r.db.
		Preload("Product_Category").
		Where("id = ?", productID).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (r *productRepo) SearchProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Preload("Product_Category").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (r *productRepo) SearchProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
	if err := r.db.Preload("Product_Category").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (r *productRepo) FilterProductByStatus(status string, product *[]pe.Product) ([]pe.Product, error) {
	if err := r.db.Where("status = ?", status).Preload("Product_Category").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

// func (r *productRepo) ImportProductFromCSV(product *pe.Product) error {
// 	if err := r.db.Save(&product).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
