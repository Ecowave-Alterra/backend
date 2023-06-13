package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (rr *reviewRepo) GetAllProducts(products *[]pe.Product) ([]pe.Product, error) {
	if err := rr.db.Preload("ProductCategory").Find(&products).Error; err != nil {
		return nil, err
	}

	return *products, nil
}

func (rr *reviewRepo) GetAllTransactionDetails(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error) {
	if err := rr.db.Where("product_id = ?", productID).Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return *transactionDetails, nil
}
