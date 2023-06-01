package common

import (
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	userHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user"
)

type Handler struct {
	ProductHandler *productHandler.ProductHandler
	UserHandler    *userHandler.UserHandler
}
