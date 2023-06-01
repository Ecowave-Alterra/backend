package information

import (
	ui "github.com/berrylradianh/ecowave-go/modules/usecase/information"
)

type InformationHandler struct {
	informationUsecase ui.InformationUsecase
}

func New(informationUsecase ui.InformationUsecase) *InformationHandler {
	return &InformationHandler{
		informationUsecase,
	}
}
