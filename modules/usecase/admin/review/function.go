package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (rc *reviewUsecase) GetAllProducts(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	return rc.reviewRepo.GetAllProducts(products, offset, pageSize)
}

func (rc *reviewUsecase) GetProductByIDNoPagination(productId string, product *pe.Product) (pe.Product, error) {
	return rc.reviewRepo.GetProductByIDNoPagination(productId, product)
}

func (rc *reviewUsecase) GetProductByID(productId string, product *pe.Product, offset, pageSize int) (*pe.Product, int64, error) {
	return rc.reviewRepo.GetProductByID(productId, product, offset, pageSize)
}

func (rc *reviewUsecase) GetProductByName(name string, product *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	return rc.reviewRepo.GetProductByName(name, product, offset, pageSize)
}

func (rc *reviewUsecase) GetAllProductByCategory(category string, product *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	return rc.reviewRepo.GetAllProductByCategory(category, product, offset, pageSize)
}

func (rc *reviewUsecase) GetAllTransactionDetailsNoPagination(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error) {
	return rc.reviewRepo.GetAllTransactionDetailsNoPagination(productID, transactionDetails)
}

func (rc *reviewUsecase) GetAllTransactionDetail(productID string, transactionDetails *[]te.TransactionDetail, offset, pageSize int) ([]te.TransactionDetail, int64, error) {
	return rc.reviewRepo.GetAllTransactionDetail(productID, transactionDetails, offset, pageSize)
}

func (rc reviewUsecase) GetTransactionByID(transactionID uint, transaction *te.Transaction) (te.Transaction, error) {
	return rc.reviewRepo.GetTransactionByID(transactionID, transaction)
}

func (rc reviewUsecase) GetUserByID(userID string, user *ue.User) (ue.User, error) {
	return rc.reviewRepo.GetUserByID(userID, user)
}

func (rc *reviewUsecase) GetAllReviewByID(reviewID string, review *re.Review) (re.Review, error) {
	return rc.reviewRepo.GetAllReviewByID(reviewID, review)
}
