package auth

import (
	"net/http"

	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (ah *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.RegisterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := ah.authUsecase.Register(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Register Sukses",
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		user, token, err := ah.authUsecase.Login(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		if user.RoleId != 2 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Email atau password salah",
				"Status":  http.StatusUnauthorized,
			})
		}

		authResponse := ue.AuthResponse{
			ID:    int(user.ID),
			Email: user.Email,
			Token: token,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil login",
			"Data":    authResponse,
		})
	}
}

func (ah *AuthHandler) LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		user, token, err := ah.authUsecase.Login(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		if user.RoleId != 1 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Email atau password salah",
				"Status":  http.StatusUnauthorized,
			})
		}

		authResponse := ue.AuthResponse{
			ID:    int(user.ID),
			Email: user.Email,
			Token: token,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil login",
			"Data":    authResponse,
		})
	}
}
