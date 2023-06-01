package auth

import "github.com/labstack/echo/v4"

func (ah *AuthHandler) RegisterRoutes(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.POST("/login", ah.LoginAdmin)
}
