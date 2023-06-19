package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ic *informationUsecase) GetAllInformations(offset int, pageSize int) (*[]ie.UserInformationResponse, int64, error) {
	informations, total, err := ic.informationRepo.GetAllInformations(offset, pageSize)
	if err != nil {
		return informations, 0, err
	}
	return informations, total, nil
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
