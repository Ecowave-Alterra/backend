package profile

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
)

func (ph *ProfileHandler) GetUserProfile(c echo.Context) error {
	var user ut.User
	var userDetail ut.UserDetail
	var addresses []ut.UserAddress
	var addressResponses []ut.UserAddressResponse

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	// log.Println(userId)

	idUserSementara := 3

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	available, err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	if !available {
		userDetail.Name = ""
		userDetail.ProfilePhotoUrl = ""
		userDetail.Phone = ""
	}

	if err := ph.profileUsecase.GetAllAddressProfileNoPagination(&addresses, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
		})
	}

	for _, address := range addresses {
		addressResponse := ut.UserAddressResponse{
			Id:           int(address.ID),
			Recipient:    address.Recipient,
			Phone:        address.Phone,
			ProvinceId:   address.ProvinceId,
			ProvinceName: address.ProvinceName,
			CityId:       address.CityId,
			CityName:     address.CityName,
			Address:      address.Address,
			Note:         address.Note,
			Mark:         address.Mark,
			IsPrimary:    address.IsPrimary,
		}

		addressResponses = append(addressResponses, addressResponse)
	}

	userResponse := ut.UserResponse{
		Id:              int(user.ID),
		GoogleId:        user.GoogleId,
		RoleId:          int(user.RoleId),
		Name:            userDetail.Name,
		Username:        user.Username,
		Email:           user.Email,
		Phone:           userDetail.Phone,
		Point:           userDetail.Point,
		ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
		Addresses:       addressResponses,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Profil berhasil didapatkan",
		"Data":    userResponse,
		"Status":  http.StatusOK,
	})
}

func (ph *ProfileHandler) UpdateUserProfile(c echo.Context) error {
	var allUser []ut.User
	var user ut.User
	var userDetail ut.UserDetail
	var userDetailBefore ut.UserDetail

	var message string
	var messagePhoto string

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	idUserSementara := 3

	if err := ph.profileUsecase.GetAllUserProfile(&allUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	available, err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	availableBefore, err := ph.profileUsecase.GetUserDetailProfile(&userDetailBefore, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	name := c.FormValue("Name")
	email := c.FormValue("Email")
	username := c.FormValue("Username")
	phone := c.FormValue("Phone")
	fileHeader, err := c.FormFile("ProfilePhotoUrl")

	if name != "" {
		userDetail.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if username != "" {
		user.Username = username
	}
	if phone != "" {
		userDetail.Phone = phone
	}

	if fileHeader != nil {
		cloudstorage.Folder = "img/users/profile/"

		if userDetailBefore.ProfilePhotoUrl != "" {
			fileName, _ := cloudstorage.GetFileName(userDetailBefore.ProfilePhotoUrl)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mendapatkan nama file",
				})
			}
			err := cloudstorage.DeleteImage(fileName)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal menghapus file pada cloud storage",
				})
			}
		}

		profilePhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)
		if profilePhotoUrl == "" {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Ups! Foto profil gagal diunggah. Coba lagi ya",
			})
		}

		userDetail.ProfilePhotoUrl = profilePhotoUrl
		messagePhoto = "Berhasil! Foto profil berhasil diubah"
	}

	for _, value := range allUser {
		if value.Username == username {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Username sudah digunakan sebelumnya",
				"Status":  http.StatusBadRequest,
			})
		}
	}

	if err := c.Validate(user); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			log.Println(validationErr)
			for _, e := range validationErr {
				if e.Tag() == "email" {
					message = "Alamat email tidak valid"
				}
			}

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": message,
				"Status":  http.StatusInternalServerError,
			})
		}
	}

	if err := ph.profileUsecase.UpdateUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Ups! Ada kendala saat mengubah profil kamu. Coba lagi ya",
		})
	} else {
		message = "Yey! Profil kamu berhasil diubah"
	}

	if !available && !availableBefore {
		userDetail = ut.UserDetail{
			Name:            userDetail.Name,
			Phone:           userDetail.Phone,
			ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
			UserId:          uint(idUserSementara),
		}

		if err := ph.profileUsecase.CreateUserDetailProfile(&userDetail); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Gagal membuat user detail profil",
			})
		}
	}

	if err := ph.profileUsecase.UpdateUserDetailProfile(&userDetail, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Ups! Ada kendala saat mengubah profil kamu. Coba lagi ya",
		})
	} else {
		message = "Yey! Profil kamu berhasil diubah"
	}

	if name == "" && email == "" && username == "" && phone == "" {
		message = messagePhoto
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": message,
		"Status":  http.StatusOK,
	})
}

func (ph *ProfileHandler) CreateAddressProfile(c echo.Context) error {
	var address ut.UserAddress

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	address.UserId = 2

	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
		})
	}

	if err := c.Validate(address); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErr {
				if e.Tag() == "required" && e.Field() == "Recipient" {
					message = "Nama penerima wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "Phone" {
					message = "Nomor telepon wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "ProvinceId" {
					message = "Id provinsi wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "ProvinceName" {
					message = "Nama provinsi wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "CityId" {
					message = "Id kota wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "CityName" {
					message = "Nama kota wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "Address" {
					message = "Alamat lengkap wajib diisi"
				}
				if e.Tag() == "max" {
					message = "Nomor telepon tidak boleh lebih dari 13 digit"
				}
			}

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": message,
				"Status":  http.StatusInternalServerError,
			})
		}
	}

	checkPhone := ""
	for i := 0; i < len(address.Phone); i++ {
		if i == 2 {
			break
		}
		checkPhone += string(address.Phone[i])
	}

	if checkPhone != "08" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Pastikan nomor kamu dimulai dengan '08'",
			"Status":  http.StatusInternalServerError,
		})
	}

	if address.IsPrimary {
		if err := ph.profileUsecase.UpdateAddressPrimaryProfile(&address, int(address.UserId)); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Gagal mengubah alamat utama",
			})
		}
		address.IsPrimary = true
	} else {
		address.IsPrimary = false
	}

	if err := ph.profileUsecase.CreateAddressProfile(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal membuat alamat",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"Message": "Yey! Kamu berhasil menambahkan alamat",
		"Status":  http.StatusCreated,
	})
}

func (ph *ProfileHandler) GetAllAddressProfile(c echo.Context) error {
	var addresses *[]ut.UserAddress
	var addressResponses []ut.UserAddressResponse

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	idUserSementara := 2

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	addresses, total, err := ph.profileUsecase.GetAllAddressProfile(addresses, idUserSementara, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	if addresses == nil || len(*addresses) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Belum ada list alamat",
			"Status":  http.StatusNotFound,
		})
	}

	for _, address := range *addresses {
		addressResponse := ut.UserAddressResponse{
			Id:           int(address.ID),
			Recipient:    address.Recipient,
			Phone:        address.Phone,
			ProvinceId:   address.ProvinceId,
			ProvinceName: address.ProvinceName,
			CityId:       address.CityId,
			CityName:     address.CityName,
			Address:      address.Address,
			Note:         address.Note,
			Mark:         address.Mark,
			IsPrimary:    address.IsPrimary,
		}

		addressResponses = append(addressResponses, addressResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":   "Alamat berhasil didapatkan",
		"Data":      addressResponses,
		"Page":      page,
		"Status":    http.StatusOK,
		"TotalPage": totalPages,
	})
}

func (ph *ProfileHandler) UpdateAddressProfile(c echo.Context) error {
	var address ut.UserAddress

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	address.UserId = 2

	idAddress, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := ph.profileUsecase.GetAddressByIdProfile(&address, int(address.UserId), idAddress); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
		})
	}

	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
		})
	}

	if err := c.Validate(address); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErr {
				if e.Tag() == "max" {
					message = "Nomor telepon tidak boleh lebih dari 13 digit"
				}
			}

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": message,
				"Status":  http.StatusInternalServerError,
			})
		}
	}

	checkPhone := ""
	for i := 0; i < len(address.Phone); i++ {
		if i == 2 {
			break
		}
		checkPhone += string(address.Phone[i])
	}

	if checkPhone != "08" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Pastikan nomor kamu dimulai dengan '08'",
			"Status":  http.StatusInternalServerError,
		})
	}

	if address.IsPrimary {
		if err := ph.profileUsecase.UpdateAddressPrimaryProfile(&address, int(address.UserId)); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Gagal mengubah alamat utama",
			})
		}
		address.IsPrimary = true
	} else {
		address.IsPrimary = false
	}

	if err := ph.profileUsecase.UpdateAddressByIdProfile(&address, int(address.UserId), idAddress); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mengubah alamat",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Yey! Kamu berhasil mengubah alamat",
		"Status":  http.StatusOK,
	})
}

func (ph *ProfileHandler) UpdatePasswordProfile(c echo.Context) error {
	var user ut.User
	var userPassword ut.UserPasswordRequest

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	idUserSementara := 2

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
		})
	}

	if err := c.Bind(&userPassword); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
		})
	}

	if len(userPassword.Password) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Password harus mengandung minimal 8 karakter",
			"Status":  http.StatusBadRequest,
		})
	}

	if userPassword.Password != userPassword.ConfirmNewPassword {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Password tidak cocok",
			"Status":  http.StatusBadRequest,
		})
	}

	message, err := ph.profileUsecase.UpdatePasswordProfile(&user, userPassword.OldPassword, userPassword.Password, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": message,
			"Status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Password berhasil diubah",
		"Status":  http.StatusOK,
	})
}

func (ph *ProfileHandler) GetAllProvince(c echo.Context) error {
	provinces, err := ph.profileUsecase.GetAllProvince()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "success",
		"Province": provinces,
	})
}

func (ph *ProfileHandler) GetAllCityByProvince(c echo.Context) error {
	provinceId := c.QueryParam("province")

	cities, err := ph.profileUsecase.GetAllCityByProvince(provinceId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "success",
		"Province": cities,
	})
}
