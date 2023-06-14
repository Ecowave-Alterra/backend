package order

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type OrderRepo interface {
	GetOrder(filter string, idUser uint, offset int, pageSize int) ([]et.Transaction, int64, error)
	OrderDetail(id uint) (et.Transaction, error)
	GetNameProductandImageUrl(id uint) (string, string, error)
	GetPromoName(id uint) (string, error)
	ConfirmOrder(id uint) error
	GetStatusOrder(id uint) (string, error)
	CancelOrder(id uint, canceledReason string) error
}

type orderRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db,
	}
}
