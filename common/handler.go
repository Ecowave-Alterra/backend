package common

import (
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
)

type Handler struct {
	ProductHandler     *productHandler.ProductHandler
	InformationHandler *informationHandler.InformationHandler
}
