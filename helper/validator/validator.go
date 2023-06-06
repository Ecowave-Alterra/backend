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
				} else if e.Field() == "Phone" || e.Tag() == "min" || e.Tag() == "max" || e.Tag() == "numeric" {
					message = fmt.Sprintf("%s tidak valid", e.Field())
				}
			}
			return fmt.Errorf(message)
		}
	}
	return nil
}

func ValidateLogin(user *ue.LoginRequest) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "email" {
					message = "Email yang anda masukkan tidak valid"
				}
			}
			return fmt.Errorf(message)
		}
	}
	return nil
}
