package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationRepo *informationRepo) GetAllInformationsNoPagination() (*[]ei.Information, error) {
	var informations []ei.Information
	if err := informationRepo.db.Find(&informations).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (informationRepo *informationRepo) GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information
	var count int64
	if err := informationRepo.db.Model(&ei.Information{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.db.Offset(offset).Limit(pageSize).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (informationRepo *informationRepo) GetInformationById(informationId int) (*ei.Information, error) {
	var information ei.Information
	if err := informationRepo.db.Where("information_id = ?", informationId).First(&information).Error; err != nil {
		return nil, err
	}

	return &information, nil
}

func (informationRepo *informationRepo) CreateInformation(information *ei.Information) error {
	if err := informationRepo.db.Create(&information).Error; err != nil {
		return err
	}

	return nil
}

func (informationRepo *informationRepo) CheckInformationExists(informationId uint) (bool, error) {
	var count int64
	result := informationRepo.db.Model(&ei.Information{}).Where("information_id = ?", informationId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	exists := count > 0
	return exists, nil
}

func (informationRepo *informationRepo) UpdateInformation(informationId int, information *ei.Information) error {
	result := informationRepo.db.Model(&information).Where("information_id = ?", informationId).Updates(&information)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (informationRepo *informationRepo) DeleteInformation(informationId int) error {
	var information *ei.Information
	if err := informationRepo.db.Delete(&information, "information_id = ?", informationId).Error; err != nil {
		return err
	}

	return nil
}

func (informationRepo *informationRepo) SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information
	var count int64
	if err := informationRepo.db.Model(&ei.Information{}).Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.db.Where("information_id LIKE ?", "%"+keyword+"%")).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.db.Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.db.Where("information_id LIKE ?", "%"+keyword+"%")).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (informationRepo *informationRepo) FilterInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information

	var count int64
	if err := informationRepo.db.Model(&ei.Information{}).Where("status = ?", keyword).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.db.Where("status = ?", keyword).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}
