package order

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (orderHandler *OrderHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	orderGroup := e.Group("/user/order")
	orderGroup.Use(jwtMiddleware)
	orderGroup.GET("", orderHandler.GetOrder())
	// orderGroup.GET("/:id", orderHandler.OrderDetail())
	orderGroup.GET("/confirm/:id", orderHandler.ConfirmOrder())
	orderGroup.POST("/cancel/:id", orderHandler.CancelOrder())
	orderGroup.GET("/tracking", orderHandler.Tracking())
}
