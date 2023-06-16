package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (vr *voucherRepo) CreateVoucher(voucher *ve.VoucherRequest) error {
	if err := vr.db.Create(voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) CheckVoucherExists(voucherId string) (bool, error) {
	var count int64
	result := vr.db.Model(&ve.Voucher{}).Where("voucher_id = ?", voucherId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	exists := count > 0
	return exists, nil
}

func (vr *voucherRepo) GetAllVoucher(offset, pageSize int) (*[]ve.Voucher, int64, error) {
	var vouchers []ve.Voucher
	var count int64
	if err := vr.db.Model(&ve.Voucher{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := vr.db.Offset(offset).Limit(pageSize).Preload("VoucherType").Find(&vouchers).Error; err != nil {
		return nil, 0, err
	}

	return &vouchers, count, nil
}

func (vr *voucherRepo) GetVoucherById(voucherId string) (*ve.Voucher, error) {
	var voucher ve.Voucher
	if err := vr.db.Where("voucher_id = ?", voucherId).Preload("VoucherType").First(&voucher).Error; err != nil {
		return nil, err
	}

	return &voucher, nil
}

func (vr *voucherRepo) UpdateVoucher(voucherId string, voucher *ve.Voucher) error {
	if err := vr.db.Model(&ve.Voucher{}).Where("voucher_id = ?", voucherId).Updates(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) DeleteVoucher(voucherId string, voucher *ve.Voucher) error {
	if err := vr.db.Where("id = ?", voucherId).Delete(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) FilterVouchersByType(voucherType string, vouchers *[]ve.Voucher) ([]ve.Voucher, error) {
	if err := vr.db.Preload("VoucherType").Where("voucher_type_id IN (SELECT id FROM voucher_types WHERE type = ?)", voucherType).Find(&vouchers).Error; err != nil {
		return nil, err
	}

	return *vouchers, nil
}
