package information

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (informationRepo *informationRepo) GetAllInformations(offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information
	var count int64
	if err := informationRepo.DB.Model(&ei.Information{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Offset(offset).Limit(pageSize).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (informationRepo *informationRepo) GetInformationById(id int) (*ei.Information, error) {
	var information ei.Information
	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Where("information_id = ?", id).First(&information).Error; err != nil {
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
	query := "SELECT COUNT(*) FROM information WHERE information_id = ?"
	var count int
	err := informationRepo.DB.Exec(query, informationId).Scan(&count)
	if err != nil {
		return false, err.Error
	}

	exists := count > 0
	return exists, nil
}

func (informationRepo *informationRepo) UpdateInformation(id int, information *ei.Information) error {
	query := "UPDATE information SET title = ?, photo_content_url = ?, content = ?, view_count = ?, bookmark_count = ?, status_id = ? WHERE information_id = ?"
	result := informationRepo.DB.Exec(query, information.Title, information.PhotoContentUrl, information.Content, information.ViewCount, information.BookmarkCount, information.StatusId, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (informationRepo *informationRepo) DeleteInformation(id int) error {
	query := "DELETE FROM information WHERE information_id = ?"
	result := informationRepo.DB.Exec(query, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (informationRepo *informationRepo) SearchInformations(keyword string, offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information
	var count int64
	if err := informationRepo.DB.Model(&ei.Information{}).Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.DB.Where("information_id LIKE ?", "%"+keyword+"%")).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.DB.Where("information_id LIKE ?", "%"+keyword+"%")).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (informationRepo *informationRepo) FilterInformations(keyword, offset, pageSize int) (*[]ei.Information, int64, error) {
	var informations []ei.Information

	var count int64
	if err := informationRepo.DB.Model(&ei.Information{}).Where("status_id = ?", keyword).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Where("status_id = ?", keyword).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}
