package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"gorm.io/gorm"
)

type InformationRepository interface {
	GetAllInformations() (*[]ei.Information, error)
	GetInformationById(id int) (*ei.Information, error)
	CreateInformation(information *ei.Information) error
	UpdateInformation(id int, information *ei.Information) error
	DeleteInformation(id int) error
}

type informationRepo struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) InformationRepository {
	return &informationRepo{
		DB,
	}
}
