package ecommerce

import (
	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (er *ecommerceRepo) GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return products, count, nil
}

func (er *ecommerceRepo) GetProductByID(productId string) ([]ee.QueryResponse, error) {
	var queryResponse *[]ee.QueryResponse

	result := er.db.Raw("SELECT p.id, p.product_id, p.name, pc.category, p.stock, p.price, p.status, p.description, ud.full_name, ud.profile_photo_url, rp.rating, rp.comment, rp.comment_admin, rp.photo_url, rp.video_url FROM rating_products rp JOIN transaction_details td ON(rp.transaction_detail_id = td.id) JOIN transactions t ON(td.transaction_id = t.id) JOIN users u ON(t.user_id = u.id) JOIN user_details ud ON(u.id = ud.user_id) JOIN products p ON(td.product_id = p.id) JOIN product_categories pc ON(p.product_category_id = pc.id) WHERE p.product_id = ?", productId).Scan(&queryResponse)
	if result.Error != nil {
		return *queryResponse, echo.NewHTTPError(404, result.Error)
	}

	return *queryResponse, nil
}

func (er *ecommerceRepo) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	var productImages []ep.ProductImage
	if err := er.db.Model(&ep.ProductImage{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, echo.NewHTTPError(404, err)
	}

	return productImages, nil
}

func (er *ecommerceRepo) AvgRating(productId string) (float64, error) {
	var avgRating float64

	result := er.db.Raw("SELECT AVG(rp.rating) AS rata22 FROM rating_products rp JOIN transaction_details td ON(rp.transaction_detail_id = td.id) JOIN products p ON(td.product_id = p.id) WHERE p.product_id = ?", productId).Scan(&avgRating)
	if result.Error != nil {
		return 0, echo.NewHTTPError(404, result.Error)
	}

	return avgRating, nil
}

func (er *ecommerceRepo) FilterProductByCategory(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Preload("ProductCategory").Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return products, count, nil
}

func (er *ecommerceRepo) FilterProductByCategoryAndPriceMax(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Order("price desc").Preload("ProductCategory").Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return products, count, nil
}

func (er *ecommerceRepo) FilterProductByCategoryAndPriceMin(category string, products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Order("price asc").Preload("ProductCategory").Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return products, count, nil
}

func (er *ecommerceRepo) FilterProductByAllCategoryAndPriceMax(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Order("price desc").Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return products, count, nil
}

func (er *ecommerceRepo) FilterProductByAllCategoryAndPriceMin(products *[]ep.Product, offset, pageSize int) (*[]ep.Product, int64, error) {
	var count int64

	if err := er.db.Model(&products).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := er.db.Offset(offset).Limit(pageSize).Order("price asc").Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return products, count, nil
}
