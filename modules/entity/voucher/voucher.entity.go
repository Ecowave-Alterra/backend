package voucher

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	*gorm.Model     `json:"-"`
	ID              uint `json:"Id" gorm:"primary_key"`
	VoucherId       string
	StartDate       time.Time   `json:"startDate" form:"startDate"`
	EndDate         time.Time   `json:"endDate" form:"endDate"`
	MinimumPurchase float64     `json:"minimumPurchase" form:"minimumPurchase"`
	MaximumDiscount float64     `json:"maximumDiscount" form:"maximumDiscount"`
	DiscountPercent float64     `json:"discountPercent" form:"discountPercent"`
	ClaimableCount  uint        `json:"claimableCount" form:"claimableCount"`
	MaxClaimLimit   uint        `json:"maxClaimLimit" form:"maxClaimLimit"`
	VoucherTypeID   uint        `json:"voucherTypeID" form:"voucherTypeID"`
	VoucherType     VoucherType `gorm:"foreignKey:VoucherTypeID"`
}

// type VoucherResponse struct {
// 	VoucherId      string
// 	Type           string
// 	ClaimableCount uint
// 	StartDate      string
// 	EndDate        string
// }

type VoucherResponse struct {
	VoucherId       string  `json:"VoucherId,omitempty"`
	Type            string  `json:"Type,omitempty"`
	StartDate       string  `json:"StartDate,omitempty"`
	EndDate         string  `json:"EndDate,omitempty"`
	MinimumPurchase float64 `json:"MinimumPurchase,omitempty"`
	MaximumDiscount float64 `json:"MaximumDiscount,omitempty"`
	DiscountPercent float64 `json:"DiscountPercent,omitempty"`
	ClaimableCount  uint    `json:"ClaimableCount,omitempty"`
	MaxClaimLimit   uint    `json:"MaxClaimLimit,omitempty"`
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

// type DetailVoucherResponse struct {
// 	Type            string
// 	EndDate         time.Time
// 	PhotoUrl        string
// 	MinimumPurchase float64 `json:"omitempty"`
// }
