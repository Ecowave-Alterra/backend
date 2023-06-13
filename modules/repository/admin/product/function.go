package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pr *productRepo) CreateProduct(product *pe.Product) error {
	if err := pr.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) CheckProductExist(productId string) (bool, error) {
	var count int64
	result := pr.db.Model(&pe.Product{}).Where("product_id = ?", productId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	exists := count > 0
	return exists, nil
}

func (pr *productRepo) CreateProductImage(productImage *pe.ProductImage) error {
	if err := pr.db.Save(&productImage).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) GetAllProduct(products *[]pe.Product) ([]pe.Product, error) {
	if err := pr.db.
		Preload("ProductCategory").
		Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (pr *productRepo) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	if err := pr.db.
		Preload("ProductCategory").
		Where("product_id = ?", productId).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (pr *productRepo) GetProductImageURLById(productId string, productImage *pe.ProductImage) ([]pe.ProductImage, error) {
	var productImages []pe.ProductImage
	if err := pr.db.Model(&pe.ProductImage{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, err
	}
	return productImages, nil
}

func (pr *productRepo) UpdateProduct(productId string, req *pe.ProductRequest) error {
	if err := pr.db.Model(&pe.Product{}).Where("product_id = ?", productId).Updates(pe.Product{ProductCategoryId: req.ProductCategoryId, Name: req.Name, Price: req.Price, Status: req.Status, Description: req.Description}).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) UpdateProductStock(productId string, stock uint) error {
	if err := pr.db.Exec("UPDATE products SET stock = ? WHERE product_id = ?", stock, productId).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) DeleteProduct(productId string, product *pe.Product) error {
	if err := pr.db.Where("product_id = ?", productId).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) DeleteProductImage(productID string, productImages *[]pe.ProductImage) error {
	if err := pr.db.Where("product_id = ?", productID).Delete(&productImages).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) DeleteProductImageByID(ProductImageID string, productImage *pe.ProductImage) error {
	if err := pr.db.Where("id = ?", ProductImageID).Delete(productImage).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepo) SearchProductByID(productID string, product *pe.Product) (pe.Product, error) {
	if err := pr.db.
		Preload("ProductCategory").
		Where("product_id = ?", productID).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (pr *productRepo) SearchProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
	if err := pr.db.Where("name LIKE ?", "%"+name+"%").Preload("ProductCategory").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (pr *productRepo) SearchProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
	if err := pr.db.Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (pr *productRepo) FilterProductByStatus(status string, product *[]pe.Product) ([]pe.Product, error) {
	if err := pr.db.Where("status = ?", status).Preload("ProductCategory").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}
