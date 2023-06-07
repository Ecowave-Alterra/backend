package information

import (
	"net/http"
	"strconv"

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

func (ih *InformationHandler) GetDetailInformations() echo.HandlerFunc {
	return func(e echo.Context) error {

		id := e.Param("id")
		informationDetail, err := ih.informationUsecase.GetDetailInformations(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informationDetail,
			"Status":       http.StatusOK,
		})
	}
}

func (ih *InformationHandler) UpdatePoint() echo.HandlerFunc {
	return func(e echo.Context) error {

		// user := e.Get("user").(*jwt.Token)
		// log.Println(user)
		// claims := user.Claims.(jwt.MapClaims)
		// claimsID := fmt.Sprint(claims["user_id"])
		// convClaimsID, err := strconv.Atoi(claimsID)

		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Id harus berupa angka",
				"Status":  http.StatusBadRequest,
			})
		}

		err = ih.informationUsecase.UpdatePoint(uint(id))
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
