package information

import (
	"time"

	"gorm.io/gorm"
)

type Information struct {
	*gorm.Model `json:"-"`

	InformationId   uint   `json:"InformationiId,"`
	Title           string `json:"Title," form:"Title" validate:"required,max=65"`
	PhotoContentUrl string `json:"PhotoContentUrl," form:"PhotoContentUrl"`
	Content         string `json:"Content," form:"Content" validate:"required"`
	ViewCount       uint   `json:"ViewCount," form:"ViewCount"`
	BookmarkCount   uint   `json:"BookmarkCount," form:"BookmarkCount"`
	Status          string `json:"Status" form:"Status"`
}

type UserInformationResponse struct {
	InformationId   uint
	Title           string
	PhotoContentUrl string
	Date            time.Time
}
type UserInformationDetailResponse struct {
	Title           string
	PhotoContentUrl string
	Content         string
	Date            time.Time
}
