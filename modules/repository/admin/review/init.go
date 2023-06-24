package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	GetAllProductReviews(offset, pageSize int) ([]re.GetAllReviewResponse, int64, error)
	GetAllProducts(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error)
	GetProductByIDNoPagination(productId string, product *pe.Product) (pe.Product, error)
	GetProductByID(productId string, product *pe.Product, offset, pageSize int) (*pe.Product, int64, error)
	GetProductByName(name string, product *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error)
	GetAllProductByCategory(category string, product *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error)
	GetAllTransactionDetailsNoPagination(productID string, transactionDetails []te.TransactionDetail) ([]te.TransactionDetail, error)
	GetAllTransactionDetail(productID string, transactionDetails *[]te.TransactionDetail, offset, pageSize int) ([]te.TransactionDetail, int64, error)
	GetTransactionByID(transactionID uint, transaction *te.Transaction) (te.Transaction, error)
	GetUserByID(userID string, user *ue.User) (ue.User, error)
	GetAllReviewByID(reviewID string, review *re.RatingProduct) (re.RatingProduct, error)
	GetProductReviewById(productId string, offset, pageSize int) ([]re.ReviewResponse, int64, error)
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
