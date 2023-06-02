package routes

import (
	"github.com/berrylradianh/ecowave-go/common"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	"github.com/berrylradianh/ecowave-go/middleware/log"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = &vld.CustomValidator{Validator: validator.New()}
	log.LogMiddleware(e)

	handler.UserHandler.RegisterRoutes(e)

	return e
}
