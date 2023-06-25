package transaction

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (transactionHandler *TransactionHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	transactionGroup := e.Group("/user/transaction")
	transactionGroup.POST("", transactionHandler.CreateTransaction(), jwtMiddleware)
	transactionGroup.GET("/point", transactionHandler.GetPoint(), jwtMiddleware)
	transactionGroup.GET("/voucher", transactionHandler.GetVoucherUser(), jwtMiddleware)
	transactionGroup.POST("/shipping-options", transactionHandler.ShippingOptions(), jwtMiddleware)
	transactionGroup.GET("/status-payment", transactionHandler.GetPaymentStatus(), jwtMiddleware)
	transactionGroup.POST("/midtrans/notifications", transactionHandler.MidtransNotifications())
}
