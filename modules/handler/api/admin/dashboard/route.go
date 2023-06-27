package dashboard

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (dashboardHandler *DashboardHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	informationGroup := e.Group("/admin/dashboard")
	informationGroup.Use(jwtMiddleware)
	informationGroup.GET("", dashboardHandler.GetDashboard())
}
