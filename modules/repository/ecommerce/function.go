package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (er *ecommerceRepo) GetAllProduct(products *[]ep.Product) ([]ep.Product, error) {
	if err := er.db.Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (er *ecommerceRepo) GetProductByID(productId string) ([]ee.QueryResponse, error) {
	var queryResponse *[]ee.QueryResponse

	result := er.db.Raw("SELECT p.id, p.name, pc.category, p.stock, p.price, p.status, p.description, ud.full_name, rp.rating, rp.comment, rp.comment_admin, rp.photo_url, rp.video_url FROM rating_products rp JOIN transaction_details td ON(rp.id = td.rating_product_id) JOIN transactions t ON(td.transaction_id = t.id) JOIN users u ON(t.user_id = u.id) JOIN user_details ud ON(u.id = ud.user_id) JOIN products p ON(td.producttt_id = p.id) JOIN product_categories pc ON(p.product_category_id = pc.id) WHERE p.product_id = ?", productId).Scan(&queryResponse)
	if result.Error != nil {
		return *queryResponse, result.Error
	}

	return *queryResponse, nil
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
