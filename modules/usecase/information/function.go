package information

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationUsecase *informationUsecase) GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error) {
	informations, count, err := informationUsecase.informationRepository.GetAllInformations(offset, pageSize)
	return informations, count, err
}

func (informationUsecase *informationUsecase) GetInformationById(id int) (*ei.Information, error) {
	information, err := informationUsecase.informationRepository.GetInformationById(id)
	return information, err
}

func (informationUsecase *informationUsecase) CreateInformation(information *ei.Information) error {
	for {
		informationId := randomid.GenerateRandomNumber()

		exists, err := informationUsecase.informationRepository.CheckInformationExists(informationId)
		if err != nil {
			return err
		}
		if !exists {
			information.InformationId = informationId
			break
		}
	}
	information.InformationId = randomid.GenerateRandomNumber()
	err := informationUsecase.informationRepository.CreateInformation(information)
	return err
}

func (informationUsecase *informationUsecase) UpdateInformation(id int, information *ei.Information) error {
	result := informationUsecase.informationRepository.UpdateInformation(id, information)
	return result
}

func (informationUsecase *informationUsecase) DeleteInformation(id int) error {
	err := informationUsecase.informationRepository.DeleteInformation(id)
	return err
}

func (informationUsecase *informationUsecase) SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	informations, count, err := informationUsecase.informationRepository.SearchInformations(keyword, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return informations, count, nil
}

func (informationUsecase *informationUsecase) FilterInformations(keyword, offset, pageSize int) (*[]ei.Information, int64, error) {
	informations, count, err := informationUsecase.informationRepository.FilterInformations(keyword, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return informations, count, nil
}
