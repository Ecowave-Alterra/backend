package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"gorm.io/gorm"
)

type InformationRepo interface {
	GetAllInformationsNoPagination() (*[]ie.Information, error)
	GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error)
	GetInformationById(informationId int) (*ie.Information, error)
	CreateInformation(information *ie.Information) error
	CheckInformationExists(informationId uint) (bool, error)
	UpdateInformation(informationId int, information *ie.Information) error
	DeleteInformation(informationId int) error
	SearchInformations(search, filter string, offset, pageSize int) (*[]ie.Information, int64, error)
}

type informationRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) InformationRepo {
	return &informationRepo{
		db,
	}
}