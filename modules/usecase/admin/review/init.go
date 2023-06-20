package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	rr "github.com/berrylradianh/ecowave-go/modules/repository/admin/review"
)

type ReviewUseCase interface {
	GetAllProducts(products *[]pe.Product) ([]pe.Product, error)
	GetProductByID(productId string, product *pe.Product) (pe.Product, error)
	GetProductByName(name string, product *[]pe.Product) ([]pe.Product, error)
	GetAllProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error)
	GetAllTransactionDetails(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error)
	GetTransactionByID(transactionID string, transaction *te.Transaction) (te.Transaction, error)
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
