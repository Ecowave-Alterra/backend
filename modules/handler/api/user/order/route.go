package order

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (orderHandler *OrderHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	orderGroup := e.Group("/user/history")
	orderGroup.Use(jwtMiddleware)
	orderGroup.GET("", orderHandler.GetOrder())
	orderGroup.GET("/:id", orderHandler.OrderDetail())
}
