package common

import (
	ah "github.com/berrylradianh/ecowave-go/modules/handler/api/admin"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
)

type Handler struct {
	AdminHundler   *ah.AdminHandler
	ProductHandler *productHandler.ProductHandler
}
