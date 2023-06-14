package transaction

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (transactionHandler *TransactionHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	transactionGroup := e.Group("/user/transaction")
	transactionGroup.Use(jwtMiddleware)
	transactionGroup.POST("", transactionHandler.CreateTransaction())
	transactionGroup.GET("/point", transactionHandler.GetPoint())
	transactionGroup.GET("/claim-voucher", transactionHandler.ClaimVoucher())
	transactionGroup.GET("/voucher", transactionHandler.GetVoucherUser())
	transactionGroup.GET("/voucher/:id", transactionHandler.DetailVoucher())
	transactionGroup.POST("/shipping-options", ShippingOptions())
	transactionGroup.POST("/midtrans/notifications", transactionHandler.MidtransNotifications())
}
