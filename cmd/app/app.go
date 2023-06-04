package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	profileHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/profile"
	userHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	profileRepo "github.com/berrylradianh/ecowave-go/modules/repository/profile"
	userRepo "github.com/berrylradianh/ecowave-go/modules/repository/user"
	productUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	profileUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/profile"
	userUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/user"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUsecase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	profileRepo := profileRepo.New(mysql.DB)
	profileUsecase := profileUsecase.New(profileRepo)
	profileHandler := profileHandler.New(profileUsecase)

	userRepo := userRepo.New(mysql.DB)
	userUsecase := userUsecase.New(userRepo)
	userHandler := userHandler.New(userUsecase)

	handler := common.Handler{
		ProductHandler: productHandler,
		ProfileHandler: profileHandler,
		UserHandler:    userHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
