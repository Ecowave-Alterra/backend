package voucher

import (
	vr "github.com/berrylradianh/ecowave-go/modules/repository/admin/voucher"
)

type VoucherUseCase interface {
}

type voucherUsecase struct {
	voucherRepo vr.VoucherRepo
}

func New(voucherRepo vr.VoucherRepo) *voucherUsecase {
	return &voucherUsecase{
		voucherRepo,
	}
}
