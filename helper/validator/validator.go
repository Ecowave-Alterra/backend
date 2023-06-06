package validator

import (
	"fmt"

	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"github.com/go-playground/validator"
)

func ValidateInformation(information *ie.Information) error {
	validate := validator.New()
	if err := validate.Struct(information); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "max" && e.Field() == "Title" {
					message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
				}
			}
			return fmt.Errorf(message)
		}
	}
	return nil
}
