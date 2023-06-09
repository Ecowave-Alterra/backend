package seed

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func CreateVoucherType() []*ve.VoucherType {
	voucherType := []*ve.VoucherType{
		{
			Type:     "Gratis Ongkir",
			PhotoURL: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
		},
		{
			Type:     "Diskon Belanja",
			PhotoURL: "https://storage.cloud.google.com/ecowave_storage/img/Sample.png",
		},
	}
	return voucherType
}
