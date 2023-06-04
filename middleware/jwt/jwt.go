package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// type JwtCustomClaims struct {
// 	User_Id int
// 	Email   string
// 	Exp     time.Time
// 	jwt.RegisteredClaims
// }

func CreateToken(userId int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

// func GetClaims(c echo.Context) *JwtCustomClaims {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*JwtCustomClaims)
// 	return claims
// }

func GetClaims(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims

	// user := e.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// claimsID := fmt.Sprint(claims["user_id"])
	// convClaimsID, err := strconv.Atoi(claimsID)
}

// func DeleteClaims(c echo.Context) *JwtCustomClaims {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*jwt.MapClaims)
// 	delete(claims, "user_id")
// 	delete(claims, "email")
// 	claims["exp"] = time.Now().Unix()
// 	return claims
// }
