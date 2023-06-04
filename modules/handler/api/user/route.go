package user

import (
	"github.com/labstack/echo/v4"
)

func (uh *UserHandler) RegisterRoute(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	profileGroup := e.Group("/user/login")
	// profileGroup.Use(jwtMiddleware)

	profileGroup.POST("", uh.LoginCustomer)
}
