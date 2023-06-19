package user

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type UserAddress struct {
	*gorm.Model  `json:"-"`
	Recipient    string           `json:"Recipient" form:"Recipient" validate:"required"`
	Phone        string           `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	ProvinceId   string           `json:"ProvinceId" form:"ProvinceId" validate:"required"`
	ProvinceName string           `json:"ProvinceName" form:"ProvinceName" validate:"required"`
	CityId       string           `json:"CityId" form:"CityId" validate:"required"`
	CityName     string           `json:"CityName" form:"CityName" validate:"required"`
	Address      string           `json:"Address" form:"Address" validate:"required"`
	Note         string           `json:"Note" form:"Note"`
	Mark         string           `json:"Mark" form:"Mark"`
	IsPrimary    bool             `json:"IsPrimary" form:"IsPrimary"`
	UserId       uint             `json:"UserId" form:"UserId"`
	Transactions []et.Transaction `gorm:"foreignKey:AddressId"`
}

type UserAddressResponse struct {
	Id           int
	Recipient    string
	Phone        string
	ProvinceId   string
	ProvinceName string
	CityId       string
	CityName     string
	Address      string
	Note         string
	Mark         string
	IsPrimary    bool
}
