package dashboard

import (
	"net/http"
	"strings"

	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"

	"github.com/labstack/echo/v4"
)

func (dh *DashboardHandler) GetDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		var chartData *[]de.ChartResponse
		var err error

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

		switch strings.ToLower(filter) {
		case "tahunan":
			chartData, err = dh.dashboardUsecase.GetYearlyRevenue()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": err.Error(),
					"Status":  http.StatusInternalServerError,
				})
			}
		case "bulanan":
			chartData, err = dh.dashboardUsecase.GetMonthlyRevenue()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": err.Error(),
					"Status":  http.StatusInternalServerError,
				})
			}
		case "mingguan":
			chartData, err = dh.dashboardUsecase.GetWeeklyRevenue()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": err.Error(),
					"Status":  http.StatusInternalServerError,
				})
			}
		default:
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Invalid filter parameter",
				"Status":  http.StatusBadRequest,
			})
		}

		totalRevenue, totalOrder, totalUser, top3Order, top3Review, err := dh.dashboardUsecase.GetDashboard()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"TotalRevenues":     totalRevenue,
			"TotalOrders":       totalOrder,
			"TotalUsers":        totalUser,
			"FavouriteProducts": top3Order,
			"TopReviewProducts": top3Review,
			"ChartData":         chartData,
			"Status":            http.StatusOK,
		})
	}
}
