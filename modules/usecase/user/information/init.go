package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	ir "github.com/berrylradianh/ecowave-go/modules/repository/user/information"
)

type InformationUsecase interface {
	GetAllInformations(offset int, pageSize int) (*[]ie.UserInformationResponse, int64, error)
	UpdatePoint(id uint) error
}

type informationUsecase struct {
	informationRepo ir.InformationRepo
}

func New(informationRepo ir.InformationRepo) *informationUsecase {
	return &informationUsecase{
		informationRepo,
	}
}
