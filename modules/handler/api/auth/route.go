package auth

import "github.com/labstack/echo/v4"

func (ah *AuthHandler) RegisterRoutes(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.POST("/login", ah.LoginAdmin())

	userGroup := e.Group("/user")
	userGroup.POST("/register", ah.Register())
	userGroup.POST("/register-google", ah.RegisterGoogle())
	userGroup.POST("/login", ah.LoginUser())
	userGroup.POST("/login-google", ah.LoginGoogle())
	userGroup.POST("/forgot-password", ah.ForgotPassword())
	userGroup.POST("/verifikasi-otp", ah.VerifOtp())
	userGroup.PUT("/change-password", ah.ChangePassword())
}
