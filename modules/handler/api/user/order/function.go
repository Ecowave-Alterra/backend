package order

import (
	"math"
	"net/http"
	"strconv"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"
	eo "github.com/berrylradianh/ecowave-go/modules/entity/order"

	"github.com/labstack/echo/v4"
)

func (oh *OrderHandler) GetOrder() echo.HandlerFunc {
	return func(e echo.Context) error {
		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		idUser, _ := h.GetIdUser(e)

		filter := e.QueryParam("filter")
		order, total, err := oh.orderUsecase.GetOrder(filter, idUser, offset, pageSize)

		if err != nil {

			return e.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if total != 0 {
			if page > totalPages {
				return e.JSON(http.StatusNotFound, echo.Map{
					"Message": "Halaman Tidak Ditemukan",
					"Status":  http.StatusNotFound,
				})
			}
		} else {
			page = 0
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Status":  200,
				"Message": "Succes get order",
				"Order":   order,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Status":    200,
			"Message":   "Succes get order",
			"Page":      page,
			"TotalPage": totalPages,
			"Order":     order,
		})
	}
}

func (oh *OrderHandler) Tracking() echo.HandlerFunc {
	return func(c echo.Context) error {

		resi := c.QueryParam("no")
		courier := c.QueryParam("cou")
		res, err := oh.orderUsecase.Tracking(resi, courier)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":   200,
			"Message":  "Success",
			"Tracking": res,
		})
	}

}
func (oh *OrderHandler) ConfirmOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		co := eo.ConfirmOrder{}
		c.Bind(&co)

		err := oh.orderUsecase.ConfirmOrder(co)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Confirm Order",
		})
	}

}
func (oh *OrderHandler) CancelOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		cancelOrder := eo.CanceledOrder{}
		c.Bind(&cancelOrder)

		err := oh.orderUsecase.CancelOrder(cancelOrder)
		if err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Cancel Order",
		})
	}

}
