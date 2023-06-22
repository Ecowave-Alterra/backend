package user

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type UserAddress struct {
	*gorm.Model  `json:"-"`
	UserId       uint             `json:"UserId" form:"UserId"`
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
	Transactions []et.Transaction `gorm:"foreignKey:AddressId"`
}

type UserAddressRequest struct {
	Recipient    string `json:"Recipient" form:"Recipient"`
	Phone        string `json:"Phone" form:"Phone" validate:"min=10,max=13"`
	ProvinceId   string `json:"ProvinceId" form:"ProvinceId"`
	ProvinceName string `json:"ProvinceName" form:"ProvinceName"`
	CityId       string `json:"CityId" form:"CityId"`
	CityName     string `json:"CityName" form:"CityName"`
	Address      string `json:"Address" form:"Address"`
	Note         string `json:"Note" form:"Note"`
	Mark         string `json:"Mark" form:"Mark"`
	IsPrimary    bool   `json:"IsPrimary" form:"IsPrimary"`
	UserId       uint   `json:"UserId" form:"UserId"`
}

type UserAddressResponse struct {
	Id           uint
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

type ProvinceResponse struct {
	RajaOngkir struct {
		Results []struct {
			ProvinceId string `json:"province_id"`
			Province   string `json:"province"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type Province struct {
	ProvinceId   string `json:"ProvinceId"`
	ProvinceName string `json:"ProvinceName"`
}

type CityResponse struct {
	RajaOngkir struct {
		Results []struct {
			CityId   string `json:"city_id"`
			CityName string `json:"city_name"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type City struct {
	CityId   string `json:"CityId"`
	CityName string `json:"CityName"`
}
