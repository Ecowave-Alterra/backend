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
	transactionGroup.GET("/voucher", transactionHandler.GetVoucherUser())
	transactionGroup.POST("/shipping-options", transactionHandler.ShippingOptions())
	transactionGroup.GET("/status-payment", transactionHandler.GetPaymentStatus())
	transactionGroup.POST("/midtrans/notifications", transactionHandler.MidtransNotifications())
}
