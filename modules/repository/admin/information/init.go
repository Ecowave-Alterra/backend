package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"gorm.io/gorm"
)

type InformationRepo interface {
	GetAllInformationsNoPagination() (*[]ie.Information, error)
	GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error)
	GetInformationById(informationId string) (*ie.Information, error)
	CreateInformation(information *ie.Information, informationDraft *ie.InformationDraftRequest) error
	CheckInformationExists(informationId string) (bool, error)
	UpdateInformation(informationId string, information *ie.Information) error
	DeleteInformation(informationId string) error
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
