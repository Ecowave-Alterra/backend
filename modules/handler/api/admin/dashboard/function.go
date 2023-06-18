package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (dh *DashboardHandler) GetDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := c.QueryParam("filter")

		validParams := map[string]bool{"filter": true}
		for param := range c.QueryParams() {
			if !validParams[param] {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"Message": "Masukkan parameter dengan benar",
					"Status":  http.StatusBadRequest,
				})
			}
		}

		totalIncome, totalOrder, totalUser, err := dh.dashboardUsecase.GetDashboard(filter)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusOK,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"TotalIncome": totalIncome,
			"TotalOrder":  totalOrder,
			"TotalUser":   totalUser,
			"Status":      http.StatusOK,
		})
	}
}
