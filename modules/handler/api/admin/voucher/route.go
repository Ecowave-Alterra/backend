package voucher

import (
	// "os"

	// echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (voucherHandler *VoucherHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	voucherGroup := e.Group("/admin/voucher")
	// voucherGroup.Use(jwtMiddleware)
	voucherGroup.POST("/", voucherHandler.CreateVoucher)
	voucherGroup.GET("/", voucherHandler.GetAllVoucher)
	voucherGroup.PUT("/:id", voucherHandler.UpdateVoucher)
	voucherGroup.DELETE("/:id", voucherHandler.DeleteVoucher)
}
