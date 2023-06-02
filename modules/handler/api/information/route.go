package information

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (informationHandler *InformationHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	informationGroup := e.Group("/admin/informations")
	informationGroup.Use(jwtMiddleware)
	informationGroup.GET("", informationHandler.GetAllInformations())
	informationGroup.GET("/:id", informationHandler.GetInformationById())
	informationGroup.POST("", informationHandler.CreateInformation())
	informationGroup.PUT("/:id", informationHandler.UpdateInformation())
	informationGroup.DELETE("/:id", informationHandler.DeleteInformation())
	informationGroup.GET("", informationHandler.SearchInformations())
	informationGroup.GET("/download-csv", informationHandler.DownloadCSVFile())
}
