package validator

import (
	"fmt"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Validation(request interface{}) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" && e.Field() == "Email" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "max" && e.Field() == "Title" {
					message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
				} else if e.Tag() == "required" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "email" {
					message = "Email yang anda masukkan tidak valid"
				} else if e.Field() == "Phone" || e.Tag() == "min" || e.Tag() == "max" || e.Tag() == "numeric" {
					message = fmt.Sprintf("%s tidak valid", e.Field())
				}
			}
			return fmt.Errorf(message)
		}
	}
	return nil
}

func StructValidator(c echo.Context, transaction et.Transaction) (string, error) {
	message := ""
	if err := c.Validate(transaction); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				if e.Tag() == "required" && e.Field() == "Email" {
					message = fmt.Sprintf("%s is required", e.Field())
					break
				} else if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required", e.Field())
					break
				} else if e.Tag() == "email" {
					message = "Invalid email address"
					break
				}
			}
		}
		return message, err
	}
	return message, nil
}
