package common

import (
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	ih "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
)

type Handler struct {
	AuthHandler        *ah.AuthHandler
	InformationHandler *ih.InformationHandler
}
