package validator

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// type CustomValidator struct {
// 	Validator *validator.Validate
// }

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.Validator.Struct(i)
// }

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
				} else if e.Tag() == "max" && e.Field() == "Phone" {
					message = "Nomor telepon tidak boleh lebih dari 13 digit"
				} else if e.Field() == "Phone" || e.Tag() == "min" || e.Tag() == "max" || e.Tag() == "numeric" {
					message = fmt.Sprintf("%s tidak valid", e.Field())
				}
			}
			return echo.NewHTTPError(422, message)
		}
	}
	return nil
}
