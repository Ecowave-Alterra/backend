package information

import (
	"github.com/labstack/echo/v4"
)

func (informationHandler *InformationHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	informationGroup := e.Group("/user/information")
	// informationGroup.Use(jwtMiddleware)
	informationGroup.GET("", informationHandler.GetAllInformations())
	informationGroup.GET("/:id", informationHandler.GetDetailInformations())
	informationGroup.POST("/point/:id", informationHandler.UpdatePoint())
}
