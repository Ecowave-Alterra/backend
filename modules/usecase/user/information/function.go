package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ic *informationUsecase) GetAllInformations() (*[]ie.UserInformationResponse, error) {
	informations, err := ic.informationRepo.GetAllInformations()
	return informations, err
}

func (ic *informationUsecase) GetDetailInformations(id string) (*ie.UserInformationDetailResponse, error) {
	informations, err := ic.informationRepo.GetDetailInformations(id)
	return informations, err
}
