package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (rr *reviewRepo) GetAllProducts(products *[]pe.Product) ([]pe.Product, error) {
	if err := rr.db.Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (rr *reviewRepo) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	if err := rr.db.
		Preload("ProductCategory").
		Where("id = ?", productId).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
	if err := rr.db.Where("name LIKE ?", "%"+name+"%").Preload("ProductCategory").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetAllProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
	if err := rr.db.Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetAllTransactionDetails(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error) {
	if err := rr.db.Where("product_id = ?", productID).Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return *transactionDetails, nil
}

func (rr *reviewRepo) GetAllReviewByID(reviewID string, review *re.Review) (re.Review, error) {
	if err := rr.db.Where("id = ?", reviewID).Find(&review).Error; err != nil {
		return *review, err
	}

	return *review, nil
}
