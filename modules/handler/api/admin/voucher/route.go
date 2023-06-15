package voucher

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (voucherHandler *VoucherHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	voucherGroup := e.Group("/admin/vouchers")
	voucherGroup.Use(jwtMiddleware)
	voucherGroup.POST("", voucherHandler.CreateVoucher)
	voucherGroup.GET("/:id", voucherHandler.GetVoucherById)
	voucherGroup.GET("", voucherHandler.GetAllVoucher)
	voucherGroup.PUT("/:id", voucherHandler.UpdateVoucher)
	voucherGroup.DELETE("/:id", voucherHandler.DeleteVoucher)
	voucherGroup.GET("/filter", voucherHandler.FilterVouchersByType)
}
