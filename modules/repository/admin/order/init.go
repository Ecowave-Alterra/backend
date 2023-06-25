package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type OrderRepo interface {
	CheckOrderExist(transactionId string) (bool, error)
	GetAllOrder(transaction *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error)
	GetOrderByID(transactionId string, transaction *te.TransactionDetailResponse) (te.TransactionDetailResponse, error)
	GetOrderProducts(transactionId string, products *[]te.TransactionProductDetailResponse) ([]te.TransactionProductDetailResponse, error)
	SearchOrder(search, filter string, offset, pageSize int) (*[]te.TransactionResponse, int64, error)
	UpdateReceiptNumber(transactionId string, receiptNumber string) error
}

type orderRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db,
	}
}
