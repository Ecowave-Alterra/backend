package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"gorm.io/gorm"
)

type InformationRepo interface {
	GetAllInformations() (*[]ie.UserInformationResponse, error)
	GetDetailInformations(id string) (*ie.UserInformationDetailResponse, error)
}

type informationRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) InformationRepo {
	return &informationRepo{
		db,
	}
}
