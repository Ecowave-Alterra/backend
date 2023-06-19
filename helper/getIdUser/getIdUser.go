package getiduser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetIdUser(c echo.Context) (uint, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	claimsID := fmt.Sprint(claims["user_id"])
	convClaimsID, err := strconv.Atoi(claimsID)
	if err != nil {
		return 0, errors.New("Need Login")
	}

	return uint(convClaimsID), nil
}
