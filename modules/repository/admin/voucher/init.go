package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	CreateVoucher(voucher *ve.VoucherRequest) error
	GetAllVoucher(offset, pageSize int) (*[]ve.Voucher, int64, error)
	GetVoucherById(voucherId string) (*ve.Voucher, error)
	CheckVoucherExists(voucherId string) (bool, error)
	UpdateVoucher(voucherId string, voucher *ve.Voucher) error
	DeleteVoucher(voucherId string, voucher *ve.Voucher) error
	FilterVoucher(filter string, offset, pageSize int) (*[]ve.Voucher, int64, error)
}

type voucherRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) VoucherRepo {
	return &voucherRepo{
		db,
	}
}
