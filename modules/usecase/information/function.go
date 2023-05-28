package information

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationUsecase *informationUsecase) GetAllInformations() (*[]ei.Information, error) {
	informations, err := informationUsecase.informationRepository.GetAllInformations()

	return informations, err
}

func (informationUsecase *informationUsecase) GetInformationById(id int) (*ei.Information, error) {
	information, err := informationUsecase.informationRepository.GetInformationById(id)
	return information, err
}

func (informationUsecase *informationUsecase) CreateInformation(information *ei.Information) error {
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

func (informationUsecase *informationUsecase) SearchInformations(keyword string) (*[]ei.Information, error) {
	informations, err := informationUsecase.informationRepository.SearchInformations(keyword)
	if err != nil {
		return nil, err
	}
	return informations, nil
}

func (informationUsecase *informationUsecase) FilterInformations(keyword int) (*[]ei.Information, error) {
	informations, err := informationUsecase.informationRepository.FilterInformations(keyword)
	if err != nil {
		return nil, err
	}
	return informations, nil
}
