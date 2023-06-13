package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (rc *reviewUsecase) GetAllProducts(products *[]pe.Product) ([]pe.Product, error) {
	return rc.reviewRepo.GetAllProducts(products)
}

func (rc *reviewUsecase) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	return rc.reviewRepo.GetProductByID(productId, product)
}

func (rc *reviewUsecase) GetProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
	return rc.reviewRepo.GetProductByName(name, product)
}

func (rc *reviewUsecase) GetAllProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
	return rc.reviewRepo.GetAllProductByCategory(category, product)
}

func (rc *reviewUsecase) GetAllTransactionDetails(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error) {
	return rc.reviewRepo.GetAllTransactionDetails(productID, transactionDetails)
}

func (rc *reviewUsecase) GetAllReviewByID(reviewID string, review *re.Review) (re.Review, error) {
	return rc.reviewRepo.GetAllReviewByID(reviewID, review)
}
