package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	vr "github.com/berrylradianh/ecowave-go/modules/repository/admin/voucher"
)

type VoucherUseCase interface {
	CreateVoucher(voucher *ve.Voucher) error
}

type voucherUsecase struct {
	voucherRepo vr.VoucherRepo
}

func New(voucherRepo vr.VoucherRepo) *voucherUsecase {
	return &voucherUsecase{
		voucherRepo,
	}
}
