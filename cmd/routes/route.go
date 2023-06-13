package routes

import (
	"github.com/berrylradianh/ecowave-go/common"
	"github.com/berrylradianh/ecowave-go/middleware/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	log.LogMiddleware(e)
	e.Use(middleware.CORS())
	log.LogMiddleware(e)

	handler.AuthHandler.RegisterRoutes(e)
	handler.InformationHandlerAdmin.RegisterRoutes(e)
	handler.InformationHandlerUser.RegisterRoutes(e)
	handler.TransactionHandlerUser.RegisterRoutes(e)
	handler.OrderHandlerUser.RegisterRoutes(e)
	handler.ProductCategoryHandler.RegisterRoutes(e)

	return e
}
