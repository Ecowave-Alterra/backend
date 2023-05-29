package information

import "gorm.io/gorm"

type Information struct {
	*gorm.Model `json:"-"`

	InformationId   uint           `json:"InformationiId,"`
	Title           string         `json:"Title," form:"Title" validate:"required,max=65"`
	PhotoContentUrl string         `json:"PhotoContentUrl," form:"PhotoContentUrl"`
	Content         string         `json:"Content," form:"Content" validate:"required"`
	ViewCount       uint           `json:"ViewCount," form:"ViewCount"`
	BookmarkCount   uint           `json:"BookmarkCount," form:"BookmarkCount"`
	StatusId        uint           `json:"StatusId," form:"StatusId" validate:"required"`
	Status          StatusResponse `gorm:"foreignKey:StatusId"`
}

type InformationResponse struct {
	*gorm.Model     `json:"-"`
	InformationId   uint   `json:"InformationiId,"`
	Title           string `json:"Title"`
	PhotoContentUrl string `json:"PhotoContentUrl"`
	Content         string `json:"Content,"`
	ViewCount       uint   `json:"ViewCount,"`
	BookmarkCount   uint   `json:"BookmarkCount,"`
	StatusId        uint   `json:"-"`
}

func (InformationResponse) TableName() string {
	return "informations"
}
