package information

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ic *informationUsecase) GetAllInformationsNoPagination() (*[]ie.Information, error) {
	informations, err := ic.informationRepo.GetAllInformationsNoPagination()
	return informations, err
}

func (ic *informationUsecase) GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error) {
	informations, count, err := ic.informationRepo.GetAllInformations(offset, pageSize)
	return informations, count, err
}

func (ic *informationUsecase) GetInformationById(informationId int) (*ie.Information, error) {
	information, err := ic.informationRepo.GetInformationById(informationId)
	return information, err
}

func (ic *informationUsecase) CreateInformation(information *ie.Information) error {
	if err := vld.ValidateInformation(information); err != nil {
		return err
	}

	for {
		informationId := randomid.GenerateRandomNumber()

		exists, err := ic.informationRepo.CheckInformationExists(informationId)
		if err != nil {
			return err
		}
		if !exists {
			information.InformationId = informationId
			break
		}
	}
	err := ic.informationRepo.CreateInformation(information)
	return err
}

func (ic *informationUsecase) UpdateInformation(informationId int, information *ie.Information) error {
	if err := vld.ValidateInformation(information); err != nil {
		return err
	}

	result := ic.informationRepo.UpdateInformation(informationId, information)
	return result
}

func (ic *informationUsecase) DeleteInformation(informationId int) error {
	err := ic.informationRepo.DeleteInformation(informationId)
	return err
}

func (ic *informationUsecase) SearchInformations(search, filter string, offset, pageSize int) (*[]ie.Information, int64, error) {
	informations, count, err := ic.informationRepo.SearchInformations(search, filter, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return informations, count, nil
}
