package user

import (
	"fmt"
	"net/http"

	userEntity "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user userEntity.UserRequest
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Fail create user",
			"Error":   err,
		})

	}

	if err := c.Validate(user); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErr {
				if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required", e.Field())
				} else if e.Tag() == "email" {
					message = "invalid email address"
				} else if e.Tag() == "min" || e.Tag() == "max" {
					message = "not valid phone number"
				}
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
			})
		}
	}

	err = h.userUC.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": "Fail create user",
			"Error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "Succes create new user",
	})
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	user := userEntity.User{}
	userLogin := userEntity.UserLogin{}

	err := c.Bind(&userLogin)

	user.Email = userLogin.Email
	user.Password = userLogin.Password

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Fail login",
			"Error":   err,
		})
	}
	if err := c.Validate(userLogin); err != nil {
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

	err, loginResponse := h.userUC.LoginUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Fail login",
			"Error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success login",
		"Data":    loginResponse,
	})
}
