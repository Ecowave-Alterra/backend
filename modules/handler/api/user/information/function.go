package information

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ih *InformationHandler) GetAllInformations() echo.HandlerFunc {
	return func(e echo.Context) error {

		informations, err := ih.informationUsecase.GetAllInformations()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informations,
		})
	}
}

func (ih *InformationHandler) GetDetailInformations() echo.HandlerFunc {
	return func(e echo.Context) error {

		id := e.Param("id")
		informationDetail, err := ih.informationUsecase.GetDetailInformations(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informationDetail,
		})
	}
}
