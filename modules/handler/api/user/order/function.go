package order

import (
	"net/http"
	"strconv"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"

	"github.com/labstack/echo/v4"
)

func (oh *OrderHandler) GetOrder() echo.HandlerFunc {
	return func(e echo.Context) error {

		idUser, err := h.GetIdUser(e)
		if err != nil {
			e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Need Login",
				"Status":  400,
			})
		}

		id := e.QueryParam("filter")
		order, err := oh.orderUsecase.GetOrder(id, idUser)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Order":  order,
			"Status": 200,
		})
	}
}

func (oh *OrderHandler) OrderDetail() echo.HandlerFunc {
	return func(e echo.Context) error {

		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid Id",
				"Status":  400,
			})
		}

		OrderDetail, err := oh.orderUsecase.OrderDetail(uint(id))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"OrderDetail": OrderDetail,
			"Status":      200,
		})
	}
}
