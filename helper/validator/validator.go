package validator

import (
	"fmt"

	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/go-playground/validator"
)

func ValidateRegister(user *ue.RegisterRequest) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "email" {
					message = "Email yang anda masukkan tidak valid"
				} else if e.Tag() == "min" {
					message = fmt.Sprintf("%s minimal %s karakter", e.Field(), e.Param())
				} else if e.Tag() == "max" {
					message = fmt.Sprintf("%s maksimal %s karakter", e.Field(), e.Param())
				} else if e.Tag() == "numeric" {
					message = fmt.Sprintf("%s harus berupa angka", e.Field())
				}
			}
			return fmt.Errorf(message)
		}
	}
	return nil
}
