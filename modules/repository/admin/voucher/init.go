package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	CreateVoucher(voucher *ve.Voucher) error
	GetAllVoucher(offset, pageSize int) (*[]ve.Voucher, int64, error)
	GetVoucherById(voucherId string) (*ve.Voucher, error)
	UpdateVoucher(voucherID string, voucher *ve.Voucher) error
	DeleteVoucher(voucherID string, voucher *ve.Voucher) error
	FilterVouchersByType(voucherType string, vouchers *[]ve.Voucher) ([]ve.Voucher, error)
}

type voucherRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) VoucherRepo {
	return &voucherRepo{
		db,
	}
}
