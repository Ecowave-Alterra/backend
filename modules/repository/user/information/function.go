package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ir *informationRepo) GetAllInformations() (*[]ie.UserInformationResponse, error) {
	var informations []ie.Information
	var informationsRes []ie.UserInformationResponse
	if err := ir.db.Where("status = ?", "Terbit").Find(&informations).Error; err != nil {
		return nil, err
	}

	for _, val := range informations {
		result := ie.UserInformationResponse{
			InformationId:   val.InformationId,
			Title:           val.Title,
			PhotoContentUrl: val.PhotoContentUrl,
			Date:            val.CreatedAt,
		}
		informationsRes = append(informationsRes, result)
	}

	return &informationsRes, nil
}

func (ir *informationRepo) GetDetailInformations(id string) (*ie.UserInformationDetailResponse, error) {
	var informations ie.Information

	if err := ir.db.Where("information_id = ?", id).First(&informations).Error; err != nil {
		return nil, err
	}

	informationDetail := ie.UserInformationDetailResponse{
		Title:           informations.Title,
		PhotoContentUrl: informations.PhotoContentUrl,
		Date:            informations.CreatedAt,
		Content:         informations.Content,
	}

	return &informationDetail, nil
}
