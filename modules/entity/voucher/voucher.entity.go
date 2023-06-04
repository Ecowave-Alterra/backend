package voucher

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	*gorm.Model `json:"-"`
	// VoucherID       uint        `json:"voucherID"`
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

type VoucherResponse struct {
	// VoucherID uint
	Type           string
	RemainingClaim uint
	StartDate      time.Time
	EndDate        time.Time
}
