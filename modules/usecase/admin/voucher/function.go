package voucher

import ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

func (vc *voucherUsecase) CreateVoucher(voucher *ve.Voucher) error {
	return vc.voucherRepo.CreateVoucher(voucher)
}
