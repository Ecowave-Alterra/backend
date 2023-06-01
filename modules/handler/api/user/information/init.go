package information

import (
	ic "github.com/berrylradianh/ecowave-go/modules/usecase/user/information"
)

type InformationHandler struct {
	informationUsecase ic.InformationUsecase
}

func New(informationUsecase ic.InformationUsecase) *InformationHandler {
	return &InformationHandler{
		informationUsecase,
	}
}
