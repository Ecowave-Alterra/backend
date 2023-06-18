package ecommerce

import (
	// echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (eh *EcommerceHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	productGroup := e.Group("/user/ecommerce")
	productGroup.GET("", eh.GetProductEcommerce)
	// productGroup.GET("/filter", eh.FilterProductByCategoryAndPrice)
}
