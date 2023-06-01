package information

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationUsecase *informationUsecase) GetAllInformationsNoPagination() (*[]ei.Information, error) {
	informations, err := informationUsecase.informationRepo.GetAllInformationsNoPagination()
	return informations, err
}

func (informationUsecase *informationUsecase) GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error) {
	informations, count, err := informationUsecase.informationRepo.GetAllInformations(offset, pageSize)
	return informations, count, err
}

func (informationUsecase *informationUsecase) GetInformationById(informationId int) (*ei.Information, error) {
	information, err := informationUsecase.informationRepo.GetInformationById(informationId)
	return information, err
}

func (informationUsecase *informationUsecase) CreateInformation(information *ei.Information) error {
	for {
		informationId := randomid.GenerateRandomNumber()

		exists, err := informationUsecase.informationRepo.CheckInformationExists(informationId)
		if err != nil {
			return err
		}
		if !exists {
			information.InformationId = informationId
			break
		}
	}
	err := informationUsecase.informationRepo.CreateInformation(information)
	return err
}

func (informationUsecase *informationUsecase) UpdateInformation(informationId int, information *ei.Information) error {
	result := informationUsecase.informationRepo.UpdateInformation(informationId, information)
	return result
}

func (informationUsecase *informationUsecase) DeleteInformation(informationId int) error {
	err := informationUsecase.informationRepo.DeleteInformation(informationId)
	return err
}

func (informationUsecase *informationUsecase) SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	informations, count, err := informationUsecase.informationRepo.SearchInformations(keyword, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return informations, count, nil
}

func (informationUsecase *informationUsecase) FilterInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	informations, count, err := informationUsecase.informationRepo.FilterInformations(keyword, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return informations, count, nil
}
