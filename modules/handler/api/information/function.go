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

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message":      "Success Get Informations",
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
			"Message":     "Success Get Information",
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
					if e.Tag() == "required" {
						message = fmt.Sprintf("%s is required", e.Field())
					} else if e.Tag() == "max" && e.Field() == "Title" {
						message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
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

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success Create Information",
		})
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
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success Update Information",
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

		err = informationHandler.informationUsecase.DeleteInformation(int(information.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success Delete Information",
		})
	}
}
