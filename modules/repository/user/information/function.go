package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
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
			Content:         val.Content,
			Date:            val.CreatedAt,
		}
		informationsRes = append(informationsRes, result)
	}

	return &informationsRes, nil
}

func (ir *informationRepo) UpdatePoint(id uint, point uint) error {

	updatePoint := eu.UserDetail{
		Point: point,
	}

	err := ir.db.Model(&eu.UserDetail{}).Where("id = ?", id).Updates(&updatePoint).Error
	if err != nil {
		return err
	}

	return nil
}

func (ir *informationRepo) GetPoint(id uint) (uint, error) {
	var userDetail eu.UserDetail

	if err := ir.db.Where("id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}
	point := userDetail.Point

	return point, nil

}
