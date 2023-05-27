package admin

import "github.com/labstack/echo/v4"

func (ah *AdminHandler) RegisterRoutes(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.POST("/login", ah.LoginAdmin)
}
