package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	vr "github.com/berrylradianh/ecowave-go/modules/repository/admin/voucher"
)

type VoucherUseCase interface {
	CreateVoucher(voucher *ve.VoucherRequest) error
	GetAllVoucher(offset, pageSize int) (*[]ve.Voucher, int64, error)
	GetVoucherById(voucherId string) (*ve.Voucher, error)
	UpdateVoucher(voucherId string, voucher *ve.Voucher) error
	DeleteVoucher(voucherId string, voucher *ve.Voucher) error
	FilterVoucher(filter string, offset, pageSize int) (*[]ve.Voucher, int64, error)
}

type voucherUsecase struct {
	voucherRepo vr.VoucherRepo
}

func New(voucherRepo vr.VoucherRepo) *voucherUsecase {
	return &voucherUsecase{
		voucherRepo,
	}
}
