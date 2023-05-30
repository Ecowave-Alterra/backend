package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
	ri "github.com/berrylradianh/ecowave-go/modules/repository/information"
)

type InformationUsecase interface {
	GetAllInformationsNoPagination() (*[]ei.Information, error)
	GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error)
	GetInformationById(informationId int) (*ei.Information, error)
	CreateInformation(information *ei.Information) error
	UpdateInformation(informationId int, information *ei.Information) error
	DeleteInformation(informationId int) error
	SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error)
	FilterInformations(keyword, offset, pageSize int) (*[]ei.Information, int64, error)
}

type informationUsecase struct {
	informationRepository ri.InformationRepository
}

func New(informationRepository ri.InformationRepository) *informationUsecase {
	return &informationUsecase{
		informationRepository,
	}
}
