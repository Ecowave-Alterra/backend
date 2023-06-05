package validator

import (
	"fmt"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
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
