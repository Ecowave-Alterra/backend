package order

import (
	"math"
	"net/http"
	"strconv"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (oh *OrderHandler) GetOrder() echo.HandlerFunc {
	return func(e echo.Context) error {
		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		idUser, _ := h.GetIdUser(e)

		id := e.QueryParam("filter")
		order, total, err := oh.orderUsecase.GetOrder(id, idUser, offset, pageSize)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}
		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return e.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Status":    200,
			"Page":      page,
			"TotalPage": totalPages,
			"Order":     order,
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

func (oh *OrderHandler) ConfirmOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": "Invalid param",
			})
		}

		err = oh.orderUsecase.ConfirmOrder(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
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

		id, err := strconv.Atoi(c.Param("id"))
		cr := et.CanceledReason{}
		c.Bind(&cr)
		canceledReason := cr.CanceledReason
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": "Invalid param",
			})
		}

		err = oh.orderUsecase.CancelOrder(uint(id), canceledReason)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Cancel Order",
		})
	}

}
