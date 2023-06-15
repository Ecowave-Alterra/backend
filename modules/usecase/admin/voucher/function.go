package voucher

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (vc *voucherUsecase) CreateVoucher(voucher *ve.VoucherRequest) error {
	if voucher.VoucherTypeID == 1 {
		voucher.MinimumPurchase = 0
		voucher.MaximumDiscount = 0
		voucher.DiscountPercent = 100
	}

	for {
		voucherId := randomid.GenerateRandomID()

		exists, err := vc.voucherRepo.CheckVoucherExists(voucherId)
		if err != nil {
			return err
		}
		if !exists {
			voucher.VoucherId = voucherId
			break
		}
	}

	if err := vld.Validation(voucher); err != nil {
		return err
	}

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
