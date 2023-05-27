package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
	ri "github.com/berrylradianh/ecowave-go/modules/repository/information"
)

type InformationUsecase interface {
	GetAllInformations() (*[]ei.Information, error)
	GetInformationById(id int) (*ei.Information, error)
	CreateInformation(information *ei.Information) error
	UpdateInformation(id int, information *ei.Information) error
	DeleteInformation(id int) error
}

type informationUsecase struct {
	informationRepository ri.InformationRepository
}

func New(informationRepository ri.InformationRepository) *informationUsecase {
	return &informationUsecase{
		informationRepository,
	}
}
