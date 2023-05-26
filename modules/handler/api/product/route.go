package product

import (
	// echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (productHandler *ProductHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	productGroup := e.Group("/products")
	// productGroup.Use(jwtMiddleware)
	productGroup.GET("", productHandler.GetAllProducts())
}
