package voucher

import ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

func (vr *voucherRepo) CreateVoucher(voucher *ve.Voucher) error {
	if err := vr.db.Create(voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) GetAllVoucher(vouchers *[]ve.Voucher) ([]ve.Voucher, error) {
	if err := vr.db.Preload("VoucherType").Find(&vouchers).Error; err != nil {
		return nil, err
	}

	return *vouchers, nil
}
