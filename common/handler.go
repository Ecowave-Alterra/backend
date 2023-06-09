package common

import (
	ecommerceHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/ecommerce"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
)

type Handler struct {
	ProductHandler   *productHandler.ProductHandler
	EcommerceHandler *ecommerceHandler.EcommerceHandler
}
