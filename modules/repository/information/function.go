package information

import (
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

func (informationRepo *informationRepo) UpdateInformation(id int, information *ei.Information) error {
	query := "UPDATE information SET title = ?, photo_content_url = ?, content = ?, view_count = ?, bookmark_count = ?, status_id = ? WHERE information_id = ?"
	result := informationRepo.DB.Exec(query, information.Title, information.PhotoContentUrl, information.Content, information.ViewCount, information.BookmarkCount, information.StatusId, id)
	if result.Error != nil {
		return result.Error
	}
	// result := informationRepo.DB.Model(&information).Where("id = ?", id).Save(&information)
	// if result.Error != nil {
	// 	return result.Error
	// }

	return nil
}

func (informationRepo *informationRepo) DeleteInformation(id int) error {
	query := "DELETE FROM information WHERE information_id = ?"
	result := informationRepo.DB.Exec(query, id)
	if result.Error != nil {
		return result.Error
	}
	// if err := informationRepo.DB.Where("information_id = ?", id).Delete(&ei.Information{}, id).Error; err != nil {
	// 	return err
	// }

	return nil
}

func (informationRepo *informationRepo) SearchInformations(keyword string) (*[]ei.Information, error) {
	var informations []ei.Information

	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Where("title LIKE ?", "%"+keyword+"%").Or(informationRepo.DB.Where("information_id LIKE ?", "%"+keyword+"%")).Find(&informations).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (informationRepo *informationRepo) FilterInformations(keyword int) (*[]ei.Information, error) {
	var informations []ei.Information

	if err := informationRepo.DB.Preload("Status", "deleted_at IS NULL").Where("status_id = ?", keyword).Find(&informations).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}
