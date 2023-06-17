package voucher

import (
	"errors"

	"github.com/berrylradianh/ecowave-go/helper/randomid"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (vc *voucherUsecase) CreateVoucher(voucher *ve.VoucherRequest) error {
	if voucher.VoucherTypeID == 1 {
		voucher.MinimumPurchase = 0
		voucher.MaximumDiscount = 0
		voucher.DiscountPercent = 100
	} else if voucher.VoucherTypeID == 2 {
		if voucher.MinimumPurchase == 0 || voucher.MaximumDiscount == 0 || voucher.DiscountPercent == 0 {
			//lint:ignore ST1005 Reason for ignoring this linter
			return errors.New("Anda gagal membuat voucher")
		}
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

func (vc *voucherUsecase) UpdateVoucher(voucherId string, voucher *ve.Voucher) error {
	if voucher.VoucherTypeID == 1 {
		voucher.MinimumPurchase = 0
		voucher.MaximumDiscount = 0
		voucher.DiscountPercent = 100
	} else if voucher.VoucherTypeID == 2 {
		if voucher.MinimumPurchase == 0 || voucher.MaximumDiscount == 0 || voucher.DiscountPercent == 0 {
			//lint:ignore ST1005 Reason for ignoring this linter
			return errors.New("Anda gagal mengubah voucher")
		}
	}

	return vc.voucherRepo.UpdateVoucher(voucherId, voucher)
}

func (vc *voucherUsecase) DeleteVoucher(voucherId string, voucher *ve.Voucher) error {
	return vc.voucherRepo.DeleteVoucher(voucherId, voucher)
}

func (vc *voucherUsecase) FilterVouchersByType(voucherType string, vouchers *[]ve.Voucher) ([]ve.Voucher, error) {
	return vc.voucherRepo.FilterVouchersByType(voucherType, vouchers)
}
