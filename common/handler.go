package common

import (
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
)

type Handler struct {
	AuthHandler        *ah.AuthHandler
	InformationHandler *informationHandler.InformationHandler
}
