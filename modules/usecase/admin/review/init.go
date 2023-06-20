package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	rr "github.com/berrylradianh/ecowave-go/modules/repository/admin/review"
)

type ReviewUseCase interface {
	GetAllProducts(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error)
	GetProductByID(productId string, product *pe.Product) (pe.Product, error)
	GetProductByName(name string, product *[]pe.Product) ([]pe.Product, error)
	GetAllProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error)
	GetAllTransactionDetailsNoPagination(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error)
	GetAllTransactionDetail(productID string, transactionDetails *[]te.TransactionDetail, offset, pageSize int) ([]te.TransactionDetail, int64, error)
	GetTransactionByID(transactionID uint, transaction *te.Transaction) (te.Transaction, error)
	GetUserByID(userID string, user *ue.User) (ue.User, error)
	GetAllReviewByID(reviewID string, review *re.Review) (re.Review, error)
}

type reviewUsecase struct {
	reviewRepo rr.ReviewRepo
}

func New(reviewRepo rr.ReviewRepo) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo,
	}
}
