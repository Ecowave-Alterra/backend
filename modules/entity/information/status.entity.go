package information

import "gorm.io/gorm"

type Status struct {
	*gorm.Model `json:"-"`

	ID                uint                  `json:"id,omitempty" gorm:"primary_key"`
	StatusInformation string                `json:"StatusInformation" form:"StatusInformation"`
	Informations      []InformationResponse `gorm:"foreignKey:StatusId"`
}

type StatusResponse struct {
	*gorm.Model       `json:"-"`
	StatusInformation string        `json:"StatusInformation" form:"StatusInformation"`
	Informations      []Information `gorm:"foreignKey:StatusId" json:"-"`
}

func (StatusResponse) TableName() string {
	return "statuses"
}
