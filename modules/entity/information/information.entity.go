package information

import "gorm.io/gorm"

type Information struct {
	*gorm.Model `json:"-"`

	ID              uint           `json:"id,omitempty" gorm:"primary_key"`
	Title           string         `json:"Title,omitempty" form:"Title" validate:"required, max=65"`
	PhotoContentUrl string         `json:"PhotoContentUrl,omitempty" form:"PhotoContentUrl" validate:"required"`
	Content         string         `json:"Content,omitempty" form:"Content" validate:"required"`
	ViewCount       uint           `json:"ViewCount,omitempty" form:"ViewCount"`
	BookmarkCount   uint           `json:"BookmarkCount,omitempty" form:"BookmarkCount"`
	StatusId        uint           `json:"StatusId,omitempty" form:"StatusId"`
	Status          StatusResponse `gorm:"foreignKey:StatusId"`
}

type InformationResponse struct {
	*gorm.Model     `json:"-"`
	Title           string `json:"Title,omitempty" form:"Title" validate:"required, max=65"`
	PhotoContentUrl string `json:"PhotoContentUrl,omitempty" form:"PhotoContentUrl" validate:"required"`
	Content         string `json:"Content,omitempty" form:"Content" validate:"required"`
	ViewCount       uint   `json:"ViewCount,omitempty" form:"ViewCount"`
	BookmarkCount   uint   `json:"BookmarkCount,omitempty" form:"BookmarkCount"`
	StatusId        uint   `json:"-"`
}

func (InformationResponse) TableName() string {
	return "informations"
}
