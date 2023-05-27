package admin

import (
	"net/http"

	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"

	"github.com/labstack/echo/v4"
)

func (ah *AdminHandler) LoginAdmin(c echo.Context) error {
	var admin at.Admin

	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	token, err := ah.adminUsecase.LoginAdmin(&admin)
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
