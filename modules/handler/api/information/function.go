package information

import (
	"fmt"
	"net/http"
	"strconv"

	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (informationHandler *InformationHandler) GetAllInformations() echo.HandlerFunc {
	return func(e echo.Context) error {
		var informations *[]ei.Information

		informations, err := informationHandler.informationUsecase.GetAllInformations()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if informations == nil || len(*informations) == 0 {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Belum ada list informasi",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informations,
		})
	}
}

func (informationHandler *InformationHandler) GetInformationById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information

		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "input id is not a number",
			})
		}

		information, err = informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Information": information,
		})
	}
}

func (informationHandler *InformationHandler) CreateInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information
		if err := e.Bind(&information); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid Request Body",
			})
		}

		if err := e.Validate(information); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "max" && e.Field() == "Title" {
						message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
					} else if e.Tag() == "required" {
						message = fmt.Sprintf("%s is required", e.Field())
					}
				}
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": message,
				})
			}
		}

		err := informationHandler.informationUsecase.CreateInformation(information)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if information.StatusId == 2 {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil menambahkan informasi ke dalam draft",
			})
		} else {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil menerbitkan informasi baru",
			})
		}
	}
}

func (informationHandler *InformationHandler) UpdateInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "input id is not a number",
			})
		}

		informationBefore, err := informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		information, err = informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		if err := e.Bind(&information); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid Request Body",
			})
		}

		if err := e.Validate(information); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "max" && e.Field() == "Title" {
						message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
					}
				}
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": message,
				})
			}
		}

		err = informationHandler.informationUsecase.UpdateInformation(int(information.ID), information)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
			})
		}

		if information.StatusId == 2 {
			if informationBefore.StatusId != information.StatusId {
				return e.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Informasi berhasil dipindahkan ke dalam draft",
				})
			}
		} else if information.StatusId == 1 {
			if informationBefore.StatusId != information.StatusId {
				return e.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Anda berhasil menerbitkan informasi baru",
				})
			}
		}
		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Anda berhasil mengubah informasi",
		})
	}
}

func (informationHandler *InformationHandler) DeleteInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Messages": "input id is not a number",
			})
		}

		information, err = informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		err = informationHandler.informationUsecase.DeleteInformation(int(information.InformationId))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Anda berhasil menghapus informasi",
		})
	}
}
