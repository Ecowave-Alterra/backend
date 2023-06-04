package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	informationHandlerAdmin "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	informationRepoAdmin "github.com/berrylradianh/ecowave-go/modules/repository/admin/information"
	informationUsecaseAdmin "github.com/berrylradianh/ecowave-go/modules/usecase/admin/information"

	informationHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	informationRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/information"
	informationUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/information"

	authHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	authRepo "github.com/berrylradianh/ecowave-go/modules/repository/auth"
	authUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/auth"

	voucherHandlerAdmin "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/voucher"
	voucherRepoAdmin "github.com/berrylradianh/ecowave-go/modules/repository/admin/voucher"
	voucherUsecaseAdmin "github.com/berrylradianh/ecowave-go/modules/usecase/admin/voucher"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	authRepo := authRepo.New(mysql.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	informationRepoAdmin := informationRepoAdmin.New(mysql.DB)
	informationUsecaseAdmin := informationUsecaseAdmin.New(informationRepoAdmin)
	informationHandlerAdmin := informationHandlerAdmin.New(informationUsecaseAdmin)

	informationRepoUser := informationRepoUser.New(mysql.DB)
	informationUsecaseUser := informationUsecaseUser.New(informationRepoUser)
	informationHandlerUser := informationHandlerUser.New(informationUsecaseUser)

	voucherRepoAdmin := voucherRepoAdmin.New(mysql.DB)
	voucherUsecaseAdmin := voucherUsecaseAdmin.New(voucherRepoAdmin)
	voucherHandlerAdmin := voucherHandlerAdmin.New(voucherUsecaseAdmin)

	handler := common.Handler{
		AuthHandler:             authHandler,
		InformationHandlerAdmin: informationHandlerAdmin,
		InformationHandlerUser:  informationHandlerUser,
		VoucherHandlerAdmin:     voucherHandlerAdmin,
	}

	router := routes.StartRoute(handler)
	return router
}
