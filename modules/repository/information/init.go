package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"gorm.io/gorm"
)

type InformationRepository interface {
	GetAllInformationsNoPagination() (*[]ei.Information, error)
	GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error)
	GetInformationById(informationId int) (*ei.Information, error)
	CreateInformation(information *ei.Information) error
	CheckInformationExists(informationId uint) (bool, error)
	UpdateInformation(informationId int, information *ei.Information) error
	DeleteInformation(informationId int) error
	SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error)
	FilterInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error)
}

type informationRepo struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) InformationRepository {
	return &informationRepo{
		DB,
	}
}
