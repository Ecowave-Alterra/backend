package profile

import (
	"math"
	"net/http"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"

	// "github.com/go-playground/validator"

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
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
		})
	}

	if err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
		})
	}

	if err := ph.profileUsecase.GetAllAddressProfileNoPagination(&addresses, idUserSementara); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
			"Status":  http.StatusNotFound,
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
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
		})
	}

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
		})
	}

	if err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
		})
	}

	if err := ph.profileUsecase.GetUserDetailProfile(&userDetailBefore, idUserSementara); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Username sudah digunakan sebelumnya",
				"Status":  http.StatusInternalServerError,
			})
		}
	}

	userRequest := ut.UserRequest{
		Email:    user.Email,
		Username: user.Username,
	}

	userDetailRequest := ut.UserDetailRequest{
		Name:            userDetail.Name,
		Phone:           userDetail.Phone,
		ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
	}

	if err := vld.Validation(userRequest); err != nil {
		return err
	}

	if err := vld.Validation(userDetailRequest); err != nil {
		return err
	}

	if err := ph.profileUsecase.UpdateUserProfile(&userRequest, idUserSementara); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Ups! Ada kendala saat mengubah profil kamu. Coba lagi ya",
			"Status":  http.StatusInternalServerError,
		})
	} else {
		message = "Yey! Profil kamu berhasil diubah"
	}

	if err := ph.profileUsecase.UpdateUserDetailProfile(&userDetailRequest, idUserSementara); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Ups! Ada kendala saat mengubah profil kamu. Coba lagi ya",
			"Status":  http.StatusInternalServerError,
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
	address.UserId = 3

	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
			"Status":  http.StatusBadRequest,
		})
	}

	if err := vld.Validation(address); err != nil {
		return err
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengubah alamat utama",
				"Status":  http.StatusInternalServerError,
			})
		}
		address.IsPrimary = true
	} else {
		address.IsPrimary = false
	}

	if err := ph.profileUsecase.CreateAddressProfile(&address); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal membuat alamat",
			"Status":  http.StatusInternalServerError,
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
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
			"Status":  http.StatusNotFound,
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
	var addressRequest ut.UserAddressRequest

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	address.UserId = 3

	idAddress, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
			"Status":  http.StatusBadRequest,
		})
	}

	if err := ph.profileUsecase.GetAddressByIdProfile(&address, int(address.UserId), idAddress); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
			"Status":  http.StatusNotFound,
		})
	}

	if err := c.Bind(&addressRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
			"Status":  http.StatusBadRequest,
		})
	}

	if addressRequest.Phone == "" {
		addressRequest.Phone = address.Phone
	}

	if err := vld.Validation(addressRequest); err != nil {
		return err
	}

	checkPhone := ""
	for i := 0; i < len(addressRequest.Phone); i++ {
		if i == 2 {
			break
		}
		checkPhone += string(addressRequest.Phone[i])
	}

	if checkPhone != "08" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Pastikan nomor kamu dimulai dengan '08'",
			"Status":  http.StatusInternalServerError,
		})
	}

	if addressRequest.IsPrimary {
		if err := ph.profileUsecase.UpdateAddressPrimaryProfile(&address, int(address.UserId)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengubah alamat utama",
				"Status":  http.StatusInternalServerError,
			})
		}
		addressRequest.IsPrimary = true
	} else {
		addressRequest.IsPrimary = false
	}

	if err := ph.profileUsecase.UpdateAddressByIdProfile(&addressRequest, int(address.UserId), idAddress); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengubah alamat",
			"Status":  http.StatusInternalServerError,
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
	idUserSementara := 3

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Gagal mendapatkan profil",
			"Status":  http.StatusNotFound,
		})
	}

	if err := c.Bind(&userPassword); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal",
			"Status":  http.StatusBadRequest,
		})
	}

	if len(userPassword.Password) < 8 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Password harus mengandung minimal 8 karakter",
			"Status":  http.StatusInternalServerError,
		})
	}

	if userPassword.Password != userPassword.ConfirmNewPassword {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Password tidak cocok",
			"Status":  http.StatusInternalServerError,
		})
	}

	message, err := ph.profileUsecase.UpdatePasswordProfile(&user, userPassword.OldPassword, userPassword.Password, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": message,
			"Status":  http.StatusInternalServerError,
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
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Success get all province",
		"Province": provinces,
	})
}

func (ph *ProfileHandler) GetAllCityByProvince(c echo.Context) error {
	provinceId := c.QueryParam("province")

	cities, err := ph.profileUsecase.GetAllCityByProvince(provinceId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "not found",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Success get all city by id province",
		"Province": cities,
	})
}
