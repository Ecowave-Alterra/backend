package voucher

import (
	"gorm.io/gorm"
)

type VoucherRepo interface {
}

type voucherRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) VoucherRepo {
	return &voucherRepo{
		db,
	}
}
