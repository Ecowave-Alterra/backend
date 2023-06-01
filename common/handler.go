package common

import (
	ih "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
)

type Handler struct {
	ProductHandler     *productHandler.ProductHandler
	InformationHandler *ih.InformationHandler
}
