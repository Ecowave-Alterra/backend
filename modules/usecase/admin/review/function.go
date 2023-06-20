package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
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

func (rc reviewUsecase) GetTransactionByID(transactionID string, transaction *te.Transaction) (te.Transaction, error) {
	return rc.reviewRepo.GetTransactionByID(transactionID, transaction)
}

func (rc reviewUsecase) GetUserByID(userID string, user *ue.User) (ue.User, error) {
	return rc.reviewRepo.GetUserByID(userID, user)
}

func (rc *reviewUsecase) GetAllReviewByID(reviewID string, review *re.Review) (re.Review, error) {
	return rc.reviewRepo.GetAllReviewByID(reviewID, review)
}
