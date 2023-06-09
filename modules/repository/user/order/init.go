package order

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type OrderRepo interface {
	GetOrder(id string, idUser uint) ([]et.Transaction, error)
	OrderDetail(id uint) (et.Transaction, error)
	GetNameProductandImageUrl(id uint) (string, string, error)
	GetPromoName(id uint) (string, error)
}

type orderRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db,
	}
}
