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

	transactionHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
	transactionRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/transaction"
	transactionUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/transaction"

	orderHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/order"
	orderRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
	orderUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/order"

	reviewHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/review"
	reviewRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/review"
	reviewUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/review"

	authHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	authRepo "github.com/berrylradianh/ecowave-go/modules/repository/auth"
	authUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/auth"
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

	transactionRepoUser := transactionRepoUser.New(mysql.DB)
	transactionUsecaseUser := transactionUsecaseUser.New(transactionRepoUser)
	transactionHandlerUser := transactionHandlerUser.New(transactionUsecaseUser)

	orderRepoUser := orderRepoUser.New(mysql.DB)
	orderUsecaseUser := orderUsecaseUser.New(orderRepoUser)
	orderHandlerUser := orderHandlerUser.New(orderUsecaseUser)

	reviewRepoUser := reviewRepoUser.New(mysql.DB)
	reviewUsecaseUser := reviewUsecaseUser.New(reviewRepoUser)
	reviewHandlerUser := reviewHandlerUser.New(reviewUsecaseUser)

	handler := common.Handler{
		AuthHandler:             authHandler,
		InformationHandlerAdmin: informationHandlerAdmin,
		InformationHandlerUser:  informationHandlerUser,
		TransactionHandlerUser:  transactionHandlerUser,
		OrderHandlerUser:        orderHandlerUser,
		ReviewHandlerUser:       reviewHandlerUser,
	}

	router := routes.StartRoute(handler)
	return router
}
