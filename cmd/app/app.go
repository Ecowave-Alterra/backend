package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
	informationRepo "github.com/berrylradianh/ecowave-go/modules/repository/information"
	informationUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/information"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	informationRepo := informationRepo.New(mysql.DB)
	informationUsecase := informationUsecase.New(informationRepo)
	informationHandler := informationHandler.New(informationUsecase)

	handler := common.Handler{
		InformationHandler: informationHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
