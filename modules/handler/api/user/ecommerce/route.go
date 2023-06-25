package ecommerce

import (
	"github.com/labstack/echo/v4"
)

func (eh *EcommerceHandler) RegisterRoutes(e *echo.Echo) {
	productGroup := e.Group("/user/ecommerce")
	productGroup.GET("", eh.GetProductEcommerce)
}
