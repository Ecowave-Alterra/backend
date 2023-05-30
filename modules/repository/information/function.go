package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationRepo *informationRepo) GetAllInformationsNoPagination() (*[]ei.Information, error) {
	var informations []ei.Information
	if err := informationRepo.DB.Find(&informations).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (informationRepo *informationRepo) GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information
	var count int64
	if err := informationRepo.DB.Model(&ei.Information{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.DB.Offset(offset).Limit(pageSize).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (informationRepo *informationRepo) GetInformationById(informationId int) (*ei.Information, error) {
	var information ei.Information
	if err := informationRepo.DB.Where("information_id = ?", informationId).First(&information).Error; err != nil {
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

func (informationRepo *informationRepo) CheckInformationExists(informationId uint) (bool, error) {
	var count int64
	result := informationRepo.DB.Model(&ei.Information{}).Where("information_id = ?", informationId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	exists := count > 0
	return exists, nil
}

func (informationRepo *informationRepo) UpdateInformation(informationId int, information *ei.Information) error {
	result := informationRepo.DB.Model(&information).Where("information_id = ?", informationId).Updates(&information)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (informationRepo *informationRepo) DeleteInformation(informationId int) error {
	var information *ei.Information
	if err := informationRepo.DB.Delete(&information, "information_id = ?", informationId).Error; err != nil {
		return err
	}

	return nil
}

func (informationRepo *informationRepo) SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information
	var count int64
	if err := informationRepo.DB.Model(&ei.Information{}).Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.DB.Where("information_id LIKE ?", "%"+keyword+"%")).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.DB.Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.DB.Where("information_id LIKE ?", "%"+keyword+"%")).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (informationRepo *informationRepo) FilterInformations(keyword, offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information

	var count int64
	if err := informationRepo.DB.Model(&ei.Information{}).Where("status = ?", keyword).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.DB.Where("status = ?", keyword).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}
