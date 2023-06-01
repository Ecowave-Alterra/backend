package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	profileHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/profile"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	profileRepo "github.com/berrylradianh/ecowave-go/modules/repository/profile"
	productUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	profileUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/profile"
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

	handler := common.Handler{
		ProductHandler: productHandler,
		ProfileHandler: profileHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
