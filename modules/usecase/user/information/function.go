package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ic *informationUsecase) GetAllInformations() (*[]ie.UserInformationResponse, error) {
	informations, err := ic.informationRepo.GetAllInformations()
	if err != nil {
		return informations, err
	}
	return informations, nil
}

func (ic *informationUsecase) GetDetailInformations(id string) (*ie.UserInformationDetailResponse, error) {
	informations, err := ic.informationRepo.GetDetailInformations(id)

	if err != nil {
		return informations, err
	}
	return informations, nil
}

func (ic *informationUsecase) UpdatePoint(id uint) error {

	point, err := ic.informationRepo.GetPoint(id)
	if err != nil {
		return err
	}

	point += 5

	err = ic.informationRepo.UpdatePoint(id, point)
	if err != nil {
		return err
	}

	return nil
}
