package user

import (
	"github.com/labstack/echo/v4"
)

func (userHandler *UserHandler) RegisterRoutes(e *echo.Echo) {

	userGroup := e.Group("/users")
	userGroup.POST("", userHandler.CreateUser)
	userGroup.POST("/login", userHandler.LoginUser)
}
