package transaction

import (
	"net/http"
	"strconv"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := h.GetIdUser(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  400,
				"Message": "Need Login",
			})
		}

		transaction := et.Transaction{}
		c.Bind(&transaction)
		transaction.UserId = uint(id)

		res, err := th.transactionUsecase.CreateTransaction(&transaction)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":      200,
			"Message":     "Success Create Transaction",
			"Transaction": res,
		})
	}

}

func (th *TransactionHandler) GetPoint() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := h.GetIdUser(c)

		res, err := th.transactionUsecase.GetPoint(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Get Point",
			"Point":   res,
		})
	}

}

func (th *TransactionHandler) GetVoucherUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := h.GetIdUser(c)

		res, err := th.transactionUsecase.GetVoucherUser(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Get Voucher User",
			"Voucher": res,
		})
	}

}
func (th *TransactionHandler) DetailVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": "Parameter tidak valid",
			})
		}

		res, err := th.transactionUsecase.DetailVoucher(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":        200,
			"Message":       "Success Get Detail Voucher",
			"DetailVoucher": res,
		})
	}

}

func (th *TransactionHandler) ClaimVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {

		idUser, _ := h.GetIdUser(c)

		idVoucher, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  400,
				"Message": "Invalid param id",
			})
		}
		shipCost, err := strconv.Atoi(c.QueryParam("ship-cost"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  400,
				"Message": "Invalid param ship cost",
			})
		}
		productCost, err := strconv.Atoi(c.QueryParam("procut-cost"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  400,
				"Message": "Invalid param product cost",
			})
		}

		res, err := th.transactionUsecase.ClaimVoucher(idUser, uint(idVoucher), float64(shipCost), float64(productCost))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Success Get Point",
			"Point":   res,
		})
	}
}
