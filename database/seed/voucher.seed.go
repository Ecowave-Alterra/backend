package seed

import (
	"time"

	"github.com/berrylradianh/ecowave-go/helper/randomid"
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func CreateVoucher() []*ve.Voucher {
	now := time.Now().UTC()
	endDate := now.AddDate(0, 1, 0)

	voucher := []*ve.Voucher{
		{
			VoucherTypeID:      2,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			MinimumPurchase:    50000,
			MaximumDiscount:    15000,
			DiscountPercent:    10,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      2,
		},
		{
			VoucherTypeID:      2,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			MinimumPurchase:    50000,
			MaximumDiscount:    15000,
			DiscountPercent:    10,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      2,
		},
		{
			VoucherTypeID:      2,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			MinimumPurchase:    50000,
			MaximumDiscount:    15000,
			DiscountPercent:    10,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      2,
		},
		{
			VoucherTypeID:      2,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			MinimumPurchase:    50000,
			MaximumDiscount:    15000,
			DiscountPercent:    10,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      2,
		},
		{
			VoucherTypeID:      2,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			MinimumPurchase:    50000,
			MaximumDiscount:    15000,
			DiscountPercent:    10,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      2,
		},
		{
			VoucherTypeID:      2,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			MinimumPurchase:    50000,
			MaximumDiscount:    15000,
			DiscountPercent:    10,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      2,
		},
		{
			VoucherTypeID:      1,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      3,
		},
		{
			VoucherTypeID:      1,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      3,
		},
		{
			VoucherTypeID:      1,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      3,
		},
		{
			VoucherTypeID:      1,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      3,
		},
		{
			VoucherTypeID:      1,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      3,
		},
		{
			VoucherTypeID:      1,
			VoucherId:          randomid.GenerateRandomID(),
			StartDate:          now,
			EndDate:            endDate,
			ClaimableUserCount: 1000,
			MaxClaimLimit:      3,
		},
	}
	return voucher
}
