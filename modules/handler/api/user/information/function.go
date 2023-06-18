package information

import (
	"net/http"

	h "github.com/berrylradianh/ecowave-go/helper/getIdUser"

	"github.com/labstack/echo/v4"
)

func (ih *InformationHandler) GetAllInformations() echo.HandlerFunc {
	return func(e echo.Context) error {

		informations, err := ih.informationUsecase.GetAllInformations()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informations,
			"Status":       http.StatusOK,
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
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Berhasil menambah point",
		})
	}
}
