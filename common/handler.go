package common

import (
	aih "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	ph "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product"
	pch "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product_category"
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	uih "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	uth "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
)

type Handler struct {
	AuthHandler             *ah.AuthHandler
	InformationHandlerAdmin *aih.InformationHandler
	InformationHandlerUser  *uih.InformationHandler
	TransactionHandlerUser  *uth.TransactionHandler
	ProductCategoryHandler  *pch.ProductCategoryHandler
	ProductHandler          *ph.ProductHandler
}
