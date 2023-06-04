package profile

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (ph *ProfileHandler) RegisterRoute(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	profileGroup := e.Group("/user")
	profileGroup.Use(jwtMiddleware)
	profileGroup.GET("", ph.GetUserProfile)

	profileGroup2 := e.Group("/user")
	profileGroup2.GET("/profile", ph.GetUser2Profile)
	profileGroup2.PUT("/profile", ph.UpdateUserProfile)
	profileGroup2.PUT("/add/profile", ph.UpdateUserProfile)
	profileGroup2.POST("/address", ph.CreateAddressProfile)
	profileGroup2.GET("/address", ph.GetAllAddressProfile)
	profileGroup2.PUT("/address/:id", ph.UpdateAddressProfile)
	profileGroup2.PUT("/password", ph.UpdatePasswordProfile)
}
