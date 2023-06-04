package seed

import (
	"time"

	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func CreateVoucher() []*ve.Voucher {
	voucher := []*ve.Voucher{
		{
			VoucherTypeID:   2,
			StartDate:       time.Date(2023, time.June, 04, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2023, time.July, 04, 0, 0, 0, 0, time.UTC),
			MinimumPurchase: 50000,
			MaximumDiscount: 15000,
			DiscountPercent: 10,
			ClaimableCount:  1000,
			MaxClaimLimit:   2,
		},
		{
			VoucherTypeID:  1,
			StartDate:      time.Date(2023, time.June, 21, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2023, time.July, 21, 0, 0, 0, 0, time.UTC),
			ClaimableCount: 1000,
			MaxClaimLimit:  3,
		},
	}
	return voucher
}
