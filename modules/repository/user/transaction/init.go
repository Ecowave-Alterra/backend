package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transaction *et.Transaction) error
	GetPoint(id uint) (uint, error)
	GetStock(id string) (uint, error)
	CountVoucherUser(idUser uint, idVoucher uint) (uint, error)
	GetVoucherUser(id uint, offset int, pageSize int) ([]ev.Voucher, int64, error)
	UpdateTransaction(transaction et.Transaction) error
	GetPaymentStatus(id string) (string, error)
	GetUserById(id uint) (*ue.User, error)
	UpdatePoint(id uint, point uint) error
}

type transactionRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TransactionRepo {
	return &transactionRepo{
		db,
	}
}
