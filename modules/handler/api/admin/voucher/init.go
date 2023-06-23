package voucher

import (
	vc "github.com/berrylradianh/ecowave-go/modules/usecase/admin/voucher"
)

type VoucherHandler struct {
	voucherUsecase vc.VoucherUseCase
}

func New(voucherUsecase vc.VoucherUseCase) *VoucherHandler {
	return &VoucherHandler{
		voucherUsecase,
	}
}
