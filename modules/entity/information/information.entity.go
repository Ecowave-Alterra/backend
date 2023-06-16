package information

import (
	"time"

	"gorm.io/gorm"
)

type Information struct {
	*gorm.Model `json:"-"`

	InformationId   string    `json:"InformationiId"`
	Title           string    `json:"Title," form:"Title" validate:"required,max=65"`
	PhotoContentUrl string    `json:"PhotoContentUrl," form:"PhotoContentUrl" validate:"required"`
	Content         string    `json:"Content," form:"Content" validate:"required"`
	CreatedAt       time.Time `json:"Date"`
	ViewCount       uint      `json:"ViewCount," form:"ViewCount"`
	BookmarkCount   uint      `json:"BookmarkCount," form:"BookmarkCount"`
	Status          string    `json:"Status" form:"Status" validate:"required"`
}

type InformationDraftRequest struct {
	*gorm.Model
	InformationId   string    `json:"InformationiId"`
	Title           string    `json:"Title," form:"Title" validate:"max=65"`
	PhotoContentUrl string    `json:"PhotoContentUrl," form:"PhotoContentUrl"`
	Content         string    `json:"Content," form:"Content"`
	CreatedAt       time.Time `json:"Date"`
	ViewCount       uint      `json:"ViewCount," form:"ViewCount"`
	BookmarkCount   uint      `json:"BookmarkCount," form:"BookmarkCount"`
	Status          string    `json:"Status" form:"Status"`
}

func (InformationDraftRequest) TableName() string {
	return "information"
}

type UserInformationResponse struct {
	InformationId   string
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
