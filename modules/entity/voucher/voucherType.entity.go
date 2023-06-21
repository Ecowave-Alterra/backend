package voucher

import "gorm.io/gorm"

type VoucherType struct {
	*gorm.Model `json:"-"`
	Type        string    `json:"type" form:"type"`
	PhotoURL    string    `json:"photoURL" form:"photoURL"`
	Vouchers    []Voucher `json:"-" gorm:"foreignKey:VoucherTypeID"`
}
