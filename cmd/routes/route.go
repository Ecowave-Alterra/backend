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

	handler.AuthHandler.RegisterRoutes(e)
	handler.InformationHandlerAdmin.RegisterRoutes(e)
	handler.InformationHandlerUser.RegisterRoutes(e)
	handler.VoucherHandlerAdmin.RegisterRoutes(e)
	handler.TransactionHandlerUser.RegisterRoutes(e)
	handler.OrderHandlerUser.RegisterRoutes(e)
	handler.ProductCategoryHandler.RegisterRoutes(e)
	handler.ProductHandler.RegisterRoutes(e)
	handler.DashboardHandler.RegisterRoutes(e)
	handler.ProfileHandler.RegisterRoute(e)
	return e
}
