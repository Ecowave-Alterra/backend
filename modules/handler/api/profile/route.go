package profile

import "github.com/labstack/echo/v4"

func (ph *ProfileHandler) RegisterRoute(e *echo.Echo) {
	profileGroup := e.Group("/user")
	profileGroup.GET("", ph.GetUserProfile)
	profileGroup.GET("/profile", ph.GetUser2Profile)
	profileGroup.PUT("/profile", ph.UpdateUserProfile)
	profileGroup.POST("/address", ph.CreateAddressProfile)
	profileGroup.GET("/address", ph.GetAllAddressProfile)
	profileGroup.PUT("/address/:id", ph.UpdateAddressProfile)
	profileGroup.PUT("/password", ph.UpdatePasswordProfile)
}
