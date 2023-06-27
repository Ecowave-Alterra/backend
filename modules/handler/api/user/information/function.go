package information

import (
	"math"
	"net/http"
	"strconv"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"

	"github.com/labstack/echo/v4"
)

func (ih *InformationHandler) GetAllInformations() echo.HandlerFunc {
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

		informations, total, err := ih.informationUsecase.GetAllInformations(offset, pageSize)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err,
				"Status":  http.StatusInternalServerError,
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
			"Status":       http.StatusOK,
			"Page":         page,
			"TotalPage":    totalPages,
			"Informations": informations,
		})
	}
}

func (ih *InformationHandler) UpdatePoint() echo.HandlerFunc {
	return func(e echo.Context) error {

		id, _ := h.GetIdUser(e)

		err := ih.informationUsecase.UpdatePoint(uint(id))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Berhasil menambah point",
		})
	}
}
