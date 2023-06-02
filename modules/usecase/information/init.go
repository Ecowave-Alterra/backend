package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	ir "github.com/berrylradianh/ecowave-go/modules/repository/information"
)

type InformationUsecase interface {
	GetAllInformationsNoPagination() (*[]ie.Information, error)
	GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error)
	GetInformationById(informationId int) (*ie.Information, error)
	CreateInformation(information *ie.Information) error
	UpdateInformation(informationId int, information *ie.Information) error
	DeleteInformation(informationId int) error
	SearchInformations(search, filter string, offset, pageSize int) (*[]ie.Information, int64, error)
}

type informationUsecase struct {
	informationRepo ir.InformationRepo
}

func New(informationRepo ir.InformationRepo) *informationUsecase {
	return &informationUsecase{
		informationRepo,
	}
}
