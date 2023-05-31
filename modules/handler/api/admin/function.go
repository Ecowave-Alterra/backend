package admin

import (
	"fmt"
	"net/http"

	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
)

func (ah *AdminHandler) LoginAdmin(c echo.Context) error {
	admin := &at.Admin{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	if err := c.Validate(admin); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErr {
				if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required", e.Field())
				} else if e.Tag() == "email" {
					message = "invalid email address"
				}
			}

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
			})
		}
	}

	token, err := ah.adminUsecase.LoginAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}
