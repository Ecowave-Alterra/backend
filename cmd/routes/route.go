package routes

import (
	"github.com/berrylradianh/ecowave-go/common"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = &vld.CustomValidator{Validator: validator.New()}
	e.Use(middleware.CORS())

	handler.AuthHandler.RegisterRoutes(e)
	handler.InformationHandler.RegisterRoutes(e)

	return e
}
