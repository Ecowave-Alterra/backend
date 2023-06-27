package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"github.com/labstack/echo/v4"
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
	if err := vr.db.Where("voucher_id = ?", voucherId).Delete(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) FilterVoucher(filter string, offset, pageSize int) (*[]ve.Voucher, int64, error) {
	var vouchers []ve.Voucher
	var count int64

	if err := vr.db.Model(&ve.Voucher{}).
		Where("voucher_type_id IN (?)",
			vr.db.Model(&ve.VoucherType{}).Select("id").Where("type LIKE ?", "%"+filter+"%")).
		Preload("VoucherType").
		Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := vr.db.Model(&ve.Voucher{}).
		Where("voucher_type_id IN (?)",
			vr.db.Model(&ve.VoucherType{}).Select("id").Where("type LIKE ?", "%"+filter+"%")).
		Preload("VoucherType").
		Offset(offset).Limit(pageSize).
		Find(&vouchers).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return &vouchers, count, nil
}
