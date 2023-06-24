package voucher

import (
	"time"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"

	"gorm.io/gorm"
)

type Voucher struct {
	*gorm.Model        `json:"-"`
	ID                 uint `json:"Id" gorm:"primary_key"`
	VoucherId          string
	StartDate          time.Time        `json:"StartDate" form:"StartDate"`
	EndDate            time.Time        `json:"EndDate" form:"EndDate"`
	MinimumPurchase    float64          `json:"MinimumPurchase" form:"MinimumPurchase"`
	MaximumDiscount    float64          `json:"MaximumDiscount" form:"MaximumDiscount"`
	DiscountPercent    float64          `json:"DiscountPercent" form:"DiscountPercent"`
	ClaimableUserCount uint             `json:"ClaimableUserCount" form:"ClaimableUserCount"`
	MaxClaimLimit      uint             `json:"MaxClaimLimit" form:"MaxClaimLimit"`
	VoucherTypeID      uint             `json:"VoucherTypeID" form:"VoucherTypeID"`
	VoucherType        VoucherType      `gorm:"foreignKey:VoucherTypeID"`
	Transactions       []et.Transaction `gorm:"foreignKey:VoucherId"`
}

type VoucherRequest struct {
	*gorm.Model
	VoucherId          string `validate:"required"`
	StartDate          string `validate:"required"`
	EndDate            string `validate:"required"`
	MinimumPurchase    float64
	MaximumDiscount    float64
	DiscountPercent    float64
	ClaimableUserCount uint `validate:"required"`
	MaxClaimLimit      uint `validate:"required"`
	VoucherTypeID      uint `validate:"required"`
}

func (VoucherRequest) TableName() string {
	return "vouchers"
}

type VoucherResponse struct {
	VoucherId          string  `json:"VoucherId,omitempty"`
	Type               string  `json:"Type,omitempty"`
	StartDate          string  `json:"StartDate,omitempty"`
	EndDate            string  `json:"EndDate,omitempty"`
	MinimumPurchase    float64 `json:"MinimumPurchase,omitempty"`
	MaximumDiscount    float64 `json:"MaximumDiscount,omitempty"`
	DiscountPercent    float64 `json:"DiscountPercent,omitempty"`
	ClaimableUserCount uint    `json:"ClaimableUserCount,omitempty"`
	MaxClaimLimit      uint    `json:"MaxClaimLimit,omitempty"`
}

type VoucherUserResponse struct {
	Id              uint
	Type            string
	EndDate         time.Time
	PhotoUrl        string
	MinimumPurchase float64
	UserClaim       uint    `json:"UserClaim,omitempty"`
	MaximumDiscount float64 `json:"MaximumDiscount,omitempty"`
	DiscountPercent float64 `json:"DiscountPercent,omitempty"`
}
