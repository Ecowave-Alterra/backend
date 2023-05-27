package information

import (
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
	err := informationUsecase.informationRepository.CreateInformation(information)
	return err
}

func (informationUsecase *informationUsecase) UpdateProduct(id int, information *ei.Information) error {
	result := informationUsecase.informationRepository.UpdateInformation(id, information)
	return result
}

func (informationUsecase *informationUsecase) DeleteProduct(id int) error {
	err := informationUsecase.informationRepository.DeleteInformation(id)
	return err
}
