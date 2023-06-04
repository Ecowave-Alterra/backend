package voucher

import ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

func (vr *voucherRepo) CreateVoucher(voucher *ve.Voucher) error {
	if err := vr.db.Create(voucher).Error; err != nil {
		return err
	}

	return nil
}
