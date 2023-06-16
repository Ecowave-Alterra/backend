package user

import "gorm.io/gorm"

type UserAddress struct {
	*gorm.Model  `json:"-"`
	Recipient    string `json:"Recipient" form:"Recipient" validate:"required"`
	PhoneNumber  string `json:"PhoneNumber" form:"PhoneNumber" validate:"required,min=10,max=13"`
	Address      string `json:"Address" form:"Address" validate:"required"`
	ProvinceId   string `json:"ProvinceId" form:"ProvinceId"`
	ProvinceName string `json:"ProvinceName" form:"ProvinceName"`
	CityId       string `json:"CityId" form:"CityId"`
	CityName     string `json:"CityName" form:"CityName"`
	Note         string `json:"Note" form:"Note"`
	Mark         string `json:"Mark" form:"Mark"`
	IsPrimary    bool   `json:"IsPrimary" form:"IsPrimary"`
	UserId       uint   `json:"UserId" form:"UserId"`
}

type UserAddressResponse struct {
	Id           int
	Recipient    string
	PhoneNumber  string
	Address      string
	ProvinceId   string
	ProvinceName string
	CityId       string
	CityName     string
	Note         string
	Mark         string
	IsPrimary    bool
	UserId       int
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
