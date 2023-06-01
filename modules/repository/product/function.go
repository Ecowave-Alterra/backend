package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (r *productRepo) CreateProduct(product *ep.Product) error {
	if err := r.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) CreateProductImage(productImage *ep.ProductImage) error {
	if err := r.db.Save(&productImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) GetAllProduct(products *[]ep.Product) ([]ep.Product, error) {
	if err := r.db.
		Preload("ProductCategory").
		Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (r *productRepo) GetProductByID(productId string, product *ep.Product) (ep.Product, error) {
	if err := r.db.
		Preload("ProductCategory").
		First(&product, productId).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (r *productRepo) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	var productImages []ep.ProductImage
	if err := r.db.Model(&ep.ProductImage{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, err
	}
	return productImages, nil
}

func (r *productRepo) UpdateProduct(productId string, req *ep.ProductRequest) error {
	if err := r.db.Model(&ep.Product{}).Where("id = ?", productId).Updates(ep.Product{ProductCategoryId: req.ProductCategoryId, Name: req.Name, Stock: req.Stock, Price: req.Price, Status: req.Status, Description: req.Description}).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) UpdateProductImage(productID string, productImage *ep.ProductImage) error {
	return nil
}

func (r *productRepo) DeleteProduct(productId string, product *ep.Product) error {
	if err := r.db.Delete(&product, productId).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) DeleteProductImage(productID string, productImages *[]ep.ProductImage) error {
	if err := r.db.Where("product_id = ?", productID).Delete(&productImages).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) DeleteProductImageByID(ProductImageID string, productImage *ep.ProductImage) error {
	if err := r.db.Where("id = ?", ProductImageID).Delete(productImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) SearchProductByID(productID string, product *ep.Product) (ep.Product, error) {
	if err := r.db.
		Preload("ProductCategory").
		Where("id = ?", productID).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (r *productRepo) SearchProductByName(name string, product *[]ep.Product) ([]ep.Product, error) {
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Preload("ProductCategory").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (r *productRepo) SearchProductByCategory(category string, product *[]ep.Product) ([]ep.Product, error) {
	if err := r.db.Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (r *productRepo) FilterProductByStatus(status string, product *[]ep.Product) ([]ep.Product, error) {
	if err := r.db.Where("status = ?", status).Preload("ProductCategory").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}
