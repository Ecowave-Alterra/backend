package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	CreateVoucher(voucher *ve.Voucher) error
	GetAllVoucher(vouchers *[]ve.Voucher) ([]ve.Voucher, error)
	UpdateVoucher(voucherID string, voucher *ve.Voucher) error
	DeleteVoucher(voucherID string, voucher *ve.Voucher) error
}

type voucherRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) VoucherRepo {
	return &voucherRepo{
		db,
	}
}
