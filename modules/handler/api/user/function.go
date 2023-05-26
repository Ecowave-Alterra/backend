package user

import (
	"net/http"

	userEntity "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user userEntity.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Fail create user",
			"Error":   err,
		})

	}

	err = h.userUC.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": "Fail create user",
			"Error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "Succes create new user",
	})
}

func (h *UserHandler) LoginCustomer(c echo.Context) error {
	user := userEntity.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Fail login",
			"Error":   err,
		})
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
		"Status":   "200",
		"Message":  "success login",
		"Customer": loginResponse,
	})
}
