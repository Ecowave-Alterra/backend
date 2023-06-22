package common

import (
	aih "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	oha "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/order"
	ph "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product"
	pch "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product_category"
	avh "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/voucher"
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	uih "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	ohu "github.com/berrylradianh/ecowave-go/modules/handler/api/user/order"
	uth "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
)

type Handler struct {
	AuthHandler             *ah.AuthHandler
	InformationHandlerAdmin *aih.InformationHandler
	InformationHandlerUser  *uih.InformationHandler
	VoucherHandlerAdmin     *avh.VoucherHandler
	TransactionHandlerUser  *uth.TransactionHandler
	OrderHandlerUser        *ohu.OrderHandler
	ProductCategoryHandler  *pch.ProductCategoryHandler
	ProductHandler          *ph.ProductHandler
	OrderHandlerAdmin       *oha.OrderHandlerAdmin
}
