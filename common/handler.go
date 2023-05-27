package common

import (
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	pch "github.com/berrylradianh/ecowave-go/modules/handler/api/product_category"
)

type Handler struct {
	ProductHandler         *productHandler.ProductHandler
	ProductCategoryHandler *pch.ProductCategoryHandler
}
