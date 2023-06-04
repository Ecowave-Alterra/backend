package transaction

import (
	"fmt"
	"net/http"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		Id, err := h.GetIdUser(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  "400",
				"Message": err.Error(),
			})
		}

		transaction := et.Transaction{}
		c.Bind(&transaction)
		transaction.UserId = uint(Id)

		if err := c.Validate(transaction); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "required" {
						message = fmt.Sprintf("%s is required", e.Field())
					}
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"Status":  "400",
						"Message": message,
					})
				}
			}
		}

		res, err := th.transactionUsecase.CreateTransaction(&transaction)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  "400",
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":      "200",
			"Message":     "Success Create Transaction",
			"Transaction": res,
		})
	}

}
