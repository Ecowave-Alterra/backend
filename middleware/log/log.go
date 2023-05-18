package log

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "method=${method}, uri=${uri}, status=${status}, time_custom=${time_custom}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
}
