package voucher

import "gorm.io/gorm"

type VoucherType struct {
	*gorm.Model
	Type     string    `json:"type" form:"type"`
	PhotoURL string    `json:"photoURL" form:"photoURL"`
	Vouchers []Voucher `gorm:"foreignKey:VoucherTypeID"`
}
