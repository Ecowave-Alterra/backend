package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	informationRepo "github.com/berrylradianh/ecowave-go/modules/repository/user/information"
	productUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	informationUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/user/information"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUsecase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	informationRepo := informationRepo.New(mysql.DB)
	informationUsecase := informationUsecase.New(informationRepo)
	informationHandler := informationHandler.New(informationUsecase)

	handler := common.Handler{
		ProductHandler:     productHandler,
		InformationHandler: informationHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
