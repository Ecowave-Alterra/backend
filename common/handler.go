package common

import (
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
)

type Handler struct {
	ProductHandler     *productHandler.ProductHandler
	InformationHandler *informationHandler.InformationHandler
}
