package common

import (
	aih "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	avh "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/voucher"
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	uih "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
)

type Handler struct {
	AuthHandler             *ah.AuthHandler
	InformationHandlerAdmin *aih.InformationHandler
	InformationHandlerUser  *uih.InformationHandler
	VoucherHandlerAdmin     *avh.VoucherHandler
}
