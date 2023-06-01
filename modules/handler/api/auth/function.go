package auth

import (
	"fmt"
	"net/http"

	"github.com/berrylradianh/ecowave-go/middleware/jwt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (ah *AuthHandler) LoginAdmin(c echo.Context) error {
	var loginAdmin struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	loginAdmin.Email = c.FormValue("email")
	loginAdmin.Password = c.FormValue("password")

	if err := c.Validate(loginAdmin); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
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

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
			})
		}
	}

	user, err := ah.authUsecase.LoginAdmin(loginAdmin.Email, loginAdmin.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Invalid email or password",
		})
	}

	token, err := jwt.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success login",
		"Token":   token,
	})
}
