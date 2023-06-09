package routes

import (
	"github.com/berrylradianh/ecowave-go/common"
	"github.com/berrylradianh/ecowave-go/middleware/log"
	"github.com/labstack/echo/v4"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	log.LogMiddleware(e)

	handler.ProductCategoryHandler.RegisterRoutes(e)

	return e
}
