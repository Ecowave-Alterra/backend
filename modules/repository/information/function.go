package information

import (
	"errors"

	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationRepo *informationRepo) GetAllInformations() (*[]ei.Information, error) {
	var informations []ei.Information
	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Find(&informations).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (informationRepo *informationRepo) GetInformationById(id int) (*ei.Information, error) {
	var information ei.Information
	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").First(&information, id).Error; err != nil {
		return nil, err
	}

	return &information, nil
}

func (informationRepo *informationRepo) CreateInformation(information *ei.Information) error {
	if err := informationRepo.DB.Create(&information).Error; err != nil {
		return err
	}

	return nil
}

func (informationRepo *informationRepo) UpdateInformation(id int, information *ei.Information) error {
	result := informationRepo.DB.Model(&information).Where("id = ?", id).Omit("UpdatedAt").Updates(&information)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("nothing updated")
	}

	return nil
}

func (informationRepo *informationRepo) DeleteInformation(id int) error {
	if err := informationRepo.DB.Delete(&ei.Information{}, id).Error; err != nil {
		return err
	}

	return nil
}
