package order

import (
	"math"
	"net/http"
	"strconv"

	cs "github.com/berrylradianh/ecowave-go/helper/customstatus"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

func (oh *OrderHandlerAdmin) GetAllOrder(c echo.Context) error {
	var transactions []te.TransactionResponse

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	transactions, total, err := oh.orderUseCase.GetAllOrder(&transactions, offset, pageSize)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	if len(transactions) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list transaksi",
			"Status":  http.StatusNotFound,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Orders":    transactions,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}
