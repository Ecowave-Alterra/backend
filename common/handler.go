package common

import (
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	profileHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/profile"
)

type Handler struct {
	ProductHandler *productHandler.ProductHandler
	ProfileHandler *profileHandler.ProfileHandler
}
