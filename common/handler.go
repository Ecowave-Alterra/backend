package common

import (
	aih "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	arh "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/review"
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	uih "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	uth "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
)

type Handler struct {
	AuthHandler             *ah.AuthHandler
	InformationHandlerAdmin *aih.InformationHandler
	InformationHandlerUser  *uih.InformationHandler
	TransactionHandlerUser  *uth.TransactionHandler
	ReviewHandlerAdmin      *arh.ReviewHandler
}
