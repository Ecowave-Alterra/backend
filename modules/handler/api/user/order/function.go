package order

import (
	"math"
	"net/http"
	"strconv"

	cs "github.com/berrylradianh/ecowave-go/helper/customstatus"
	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (oh *OrderHandler) GetOrder() echo.HandlerFunc {
	return func(e echo.Context) error {
		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
			})
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		idUser, _ := h.GetIdUser(e)

		filter := e.QueryParam("filter")
		order, total, err := oh.orderUsecase.GetOrder(filter, idUser, offset, pageSize)

		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return e.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
			})
		}
		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Status":    200,
			"Message":   "Succes get order",
			"Page":      page,
			"TotalPage": totalPages,
			"Order":     order,
		})
	}
}

// func (oh *OrderHandler) OrderDetail() echo.HandlerFunc {
// 	return func(e echo.Context) error {

// 		id, err := strconv.Atoi(e.Param("id"))
// 		if err != nil {
// 			e.JSON(http.StatusBadRequest, map[string]interface{}{
// 				"Status":  400,
// 				"Message": "Invalid Id",
// 			})
// 		}

// 		OrderDetail, err := oh.orderUsecase.OrderDetail(uint(id))
// 		if err != nil {
// 			code, msg := cs.CustomStatus(err.Error())
// 			return e.JSON(code, echo.Map{
// 				"Status":  code,
// 				"Message": msg,
// 			})
// 		}

// 		return e.JSON(http.StatusOK, map[string]interface{}{
// 			"Status":      200,
// 			"Message":     "Succes get order detail",
// 			"OrderDetail": OrderDetail,
// 		})
// 	}
// }

func (oh *OrderHandler) Tracking() echo.HandlerFunc {
	return func(c echo.Context) error {

		resi := c.QueryParam("no")
		courier := c.QueryParam("cou")
		res, err := oh.orderUsecase.Tracking(resi, courier)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
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

		id := c.Param("id")
		err := oh.orderUsecase.ConfirmOrder(id)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
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

		id := c.Param("id")
		cr := et.CanceledReason{}
		c.Bind(&cr)
		canceledReason := cr.CanceledReason

		err := oh.orderUsecase.CancelOrder(id, canceledReason)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Cancel Order",
		})
	}

}
