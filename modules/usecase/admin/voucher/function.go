package voucher

import ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

func (vc *voucherUsecase) CreateVoucher(voucher *ve.Voucher) error {
	return vc.voucherRepo.CreateVoucher(voucher)
}

func (vc *voucherUsecase) GetAllVoucher(offset, pageSize int) (*[]ve.Voucher, int64, error) {
	vouchers, count, err := vc.voucherRepo.GetAllVoucher(offset, pageSize)
	return vouchers, count, err
}

func (vc *voucherUsecase) GetVoucherById(voucherId string) (*ve.Voucher, error) {
	voucher, err := vc.voucherRepo.GetVoucherById(voucherId)
	return voucher, err
}

func (vc *voucherUsecase) UpdateVoucher(voucherID string, voucher *ve.Voucher) error {
	return vc.voucherRepo.UpdateVoucher(voucherID, voucher)
}

func (vc *voucherUsecase) DeleteVoucher(voucherID string, voucher *ve.Voucher) error {
	return vc.voucherRepo.DeleteVoucher(voucherID, voucher)
}

func (vc *voucherUsecase) FilterVouchersByType(voucherType string, vouchers *[]ve.Voucher) ([]ve.Voucher, error) {
	return vc.voucherRepo.FilterVouchersByType(voucherType, vouchers)
}
