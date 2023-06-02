package auth

import (
	"fmt"
	"net/http"

	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (ah *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &ue.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Username: c.FormValue("username"),
			RoleId:   2,
			UserDetail: ue.UserDetail{
				Name:  c.FormValue("name"),
				Phone: c.FormValue("phone"),
			},
		}

		if err := c.Validate(user); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "required" {
						message = fmt.Sprintf("%s is required", e.Field())
					} else if e.Tag() == "email" {
						message = "Invalid email address"
					}
				}
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": message,
				})
			}
		}

		err := ah.authUsecase.Register(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Register Successfull",
		})
	}
}

func (ah *AuthHandler) LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		user, token, err := ah.authUsecase.Login(loginAdmin.Email, loginAdmin.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Invalid email or password",
			})
		}

		if user.RoleId != 1 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Invalid email or password",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success login",
			"Token":   token,
		})
	}
}

func (ah *AuthHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginUser struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}

		loginUser.Email = c.FormValue("email")
		loginUser.Password = c.FormValue("password")

		if err := c.Validate(loginUser); err != nil {
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

		user, token, err := ah.authUsecase.Login(loginUser.Email, loginUser.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Invalid email or password",
			})
		}

		if user.RoleId != 2 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Invalid email or password",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success login",
			"Token":   token,
		})
	}
}
