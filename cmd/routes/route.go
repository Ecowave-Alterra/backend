package routes

import (
	"github.com/berrylradianh/ecowave-go/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	handler.AuthHandler.RegisterRoutes(e)
	handler.InformationHandlerAdmin.RegisterRoutes(e)
	handler.InformationHandlerUser.RegisterRoutes(e)

	return e
}
