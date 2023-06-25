package order

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (orderHandler *OrderHandlerAdmin) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	orderGroup := e.Group("/admin/orders")
	orderGroup.Use(jwtMiddleware)
	orderGroup.GET("", orderHandler.GetAllOrder)
	orderGroup.GET("/:id", orderHandler.GetOrderByID)
	orderGroup.GET("/search", orderHandler.SearchOrder)
	orderGroup.PUT("/:id", orderHandler.UpdateReceiptNumber)
}
