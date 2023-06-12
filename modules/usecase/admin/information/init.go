package information

import (
	"mime/multipart"

	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	ir "github.com/berrylradianh/ecowave-go/modules/repository/admin/information"
)

type InformationUsecase interface {
	GetAllInformationsNoPagination() (*[]ie.Information, error)
	GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error)
	GetInformationById(informationId string) (*ie.Information, error)
	CreateInformation(fileHeader *multipart.FileHeader, title, content, status string) error
	CreateInformationDraft(fileHeader *multipart.FileHeader, title, content, status string) error
	UpdateInformation(informationId string, information *ie.Information) error
	DeleteInformation(informationId string) error
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
