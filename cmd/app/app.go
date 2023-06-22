package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	authHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	authRepo "github.com/berrylradianh/ecowave-go/modules/repository/auth"
	authUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/auth"

	dashboardHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/dashboard"
	dashboardRepo "github.com/berrylradianh/ecowave-go/modules/repository/admin/dashboard"
	dashboardUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/admin/dashboard"

	informationHandlerAdmin "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	informationRepoAdmin "github.com/berrylradianh/ecowave-go/modules/repository/admin/information"
	informationUsecaseAdmin "github.com/berrylradianh/ecowave-go/modules/usecase/admin/information"

	productCategoryHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product_category"
	productCategoryRepo "github.com/berrylradianh/ecowave-go/modules/repository/admin/product_category"
	productCategoryUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/admin/product_category"

	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/admin/product"
	productUseCase "github.com/berrylradianh/ecowave-go/modules/usecase/admin/product"

	informationHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	informationRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/information"
	informationUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/information"

	voucherHandlerAdmin "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/voucher"
	voucherRepoAdmin "github.com/berrylradianh/ecowave-go/modules/repository/admin/voucher"
	voucherUsecaseAdmin "github.com/berrylradianh/ecowave-go/modules/usecase/admin/voucher"

	transactionHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
	transactionRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/transaction"
	transactionUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/transaction"

	orderHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/order"
	orderRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
	orderUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/order"

	profileHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user/profile"
	profileRepo "github.com/berrylradianh/ecowave-go/modules/repository/user/profile"
	profileUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/user/profile"

	reviewHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/review"
	reviewRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/review"
	reviewUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/review"

	ecommerceHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user/ecommerce"
	ecommerceRepo "github.com/berrylradianh/ecowave-go/modules/repository/user/ecommerce"
	ecommerceUseCase "github.com/berrylradianh/ecowave-go/modules/usecase/user/ecommerce"

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

	productCategoryRepo := productCategoryRepo.New(mysql.DB)
	productCategoryUsecase := productCategoryUsecase.New(productCategoryRepo)
	productCategoryHandler := productCategoryHandler.New(productCategoryUsecase)

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUseCase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	ecommerceRepo := ecommerceRepo.New(mysql.DB)
	ecommerceUsecase := ecommerceUseCase.New(ecommerceRepo)
	ecommerceHandler := ecommerceHandler.New(ecommerceUsecase)

	profileRepo := profileRepo.New(mysql.DB)
	profileUsecase := profileUsecase.New(profileRepo)
	profileHandler := profileHandler.New(profileUsecase)

	informationRepoUser := informationRepoUser.New(mysql.DB)
	informationUsecaseUser := informationUsecaseUser.New(informationRepoUser)
	informationHandlerUser := informationHandlerUser.New(informationUsecaseUser)

	voucherRepoAdmin := voucherRepoAdmin.New(mysql.DB)
	voucherUsecaseAdmin := voucherUsecaseAdmin.New(voucherRepoAdmin)
	voucherHandlerAdmin := voucherHandlerAdmin.New(voucherUsecaseAdmin)

	transactionRepoUser := transactionRepoUser.New(mysql.DB)
	transactionUsecaseUser := transactionUsecaseUser.New(transactionRepoUser)
	transactionHandlerUser := transactionHandlerUser.New(transactionUsecaseUser)

	orderRepoUser := orderRepoUser.New(mysql.DB)
	orderUsecaseUser := orderUsecaseUser.New(orderRepoUser)
	orderHandlerUser := orderHandlerUser.New(orderUsecaseUser)

	reviewRepoUser := reviewRepoUser.New(mysql.DB)
	reviewUsecaseUser := reviewUsecaseUser.New(reviewRepoUser)
	reviewHandlerUser := reviewHandlerUser.New(reviewUsecaseUser)

	dashboardRepo := dashboardRepo.New(mysql.DB)
	dashboardUsecase := dashboardUsecase.New(dashboardRepo)
	dashboardHandler := dashboardHandler.New(dashboardUsecase)

	handler := common.Handler{
		ProductHandler:          productHandler,
		ProfileHandler:          profileHandler,
		AuthHandler:             authHandler,
		InformationHandlerAdmin: informationHandlerAdmin,
		InformationHandlerUser:  informationHandlerUser,
		VoucherHandlerAdmin:     voucherHandlerAdmin,
		TransactionHandlerUser:  transactionHandlerUser,
		OrderHandlerUser:        orderHandlerUser,
		ReviewHandlerUser:       reviewHandlerUser,
		ProductCategoryHandler:  productCategoryHandler,
		DashboardHandler:        dashboardHandler,
		EcommerceHandler:        ecommerceHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
