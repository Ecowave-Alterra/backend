package common

import (
	dh "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/dashboard"
	aih "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	ph "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product"
	pch "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product_category"
	avh "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/voucher"
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	ecommerceHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user/ecommerce"
	uih "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	ohu "github.com/berrylradianh/ecowave-go/modules/handler/api/user/order"

	profileHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user/profile"
	urh "github.com/berrylradianh/ecowave-go/modules/handler/api/user/review"
	uth "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
)

type Handler struct {
	ProfileHandler          *profileHandler.ProfileHandler
	AuthHandler             *ah.AuthHandler
	InformationHandlerAdmin *aih.InformationHandler
	InformationHandlerUser  *uih.InformationHandler
	VoucherHandlerAdmin     *avh.VoucherHandler
	TransactionHandlerUser  *uth.TransactionHandler
	OrderHandlerUser        *ohu.OrderHandler
	ReviewHandlerUser       *urh.ReviewHandler
	ProductCategoryHandler  *pch.ProductCategoryHandler
	ProductHandler          *ph.ProductHandler
	DashboardHandler        *dh.DashboardHandler
	EcommerceHandler        *ecommerceHandler.EcommerceHandler
}
