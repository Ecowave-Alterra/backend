package profile

import (
	"log"
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

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	// log.Println(userId)

	idUserSementara := 1

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
		userDetail.FullName = ""
		userDetail.EcoPoint = 0
	}

	userResponse := ut.UserResponse{
		UserId:       int(user.ID),
		FullName:     userDetail.FullName,
		Username:     user.Username,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		EcoPoint:     userDetail.EcoPoint,
		UserDetailId: int(userDetail.ID),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Profil berhasil didapatkan",
		"Data":    userResponse,
	})
}

func (ph *ProfileHandler) GetUser2Profile(c echo.Context) error {
	var user ut.User
	var userDetail ut.UserDetail

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	idUserSementara := 1

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
		userDetail.FullName = ""
		userDetail.ProfilePhotoUrl = ""
	}

	user2Response := ut.User2Response{
		UserId:          int(user.ID),
		FullName:        userDetail.FullName,
		Username:        user.Username,
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
		UserDetailId:    int(userDetail.ID),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Profil berhasil didapatkan",
		"Data":    user2Response,
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
	idUserSementara := 1

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

	fullName := c.FormValue("FullName")
	email := c.FormValue("Email")
	username := c.FormValue("Username")
	phoneNumber := c.FormValue("PhoneNumber")
	fileHeader, err := c.FormFile("ProfilePhotoUrl")

	if fullName != "" {
		userDetail.FullName = fullName
	}
	if email != "" {
		user.Email = email
	}
	if username != "" {
		user.Username = username
	}
	if phoneNumber != "" {
		user.PhoneNumber = phoneNumber
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

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": message,
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
			FullName:        userDetail.FullName,
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

	if fullName == "" && email == "" && username == "" && phoneNumber == "" {
		message = messagePhoto
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": message,
	})
}

func (ph *ProfileHandler) CreateAddressProfile(c echo.Context) error {
	var address ut.UserAddress

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	address.UserId = 1

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
				if e.Tag() == "required" && e.Field() == "PhoneNumber" {
					message = "Nomor telepon wajib diisi"
				}
				if e.Tag() == "required" && e.Field() == "Address" {
					message = "Alamat lengkap wajib diisi"
				}
				if e.Tag() == "max" {
					message = "Nomor telepon tidak boleh lebih dari 13 digit"
				}
			}

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": message,
			})
		}
	}

	checkPhoneNumber := ""
	for i := 0; i < len(address.PhoneNumber); i++ {
		if i == 2 {
			break
		}
		checkPhoneNumber += string(address.PhoneNumber[i])
	}

	if checkPhoneNumber != "08" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Pastikan nomor kamu dimulai dengan '08'",
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Alamat berhasil dibuat",
	})
}

func (ph *ProfileHandler) GetAllAddressProfile(c echo.Context) error {
	var addresses []ut.UserAddress
	var addressResponses []ut.UserAddressResponse

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	idUserSementara := 1

	if err := ph.profileUsecase.GetAllAddressProfile(&addresses, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Gagal mendapatkan alamat",
		})
	}

	for _, address := range addresses {
		addressResponse := ut.UserAddressResponse{
			UserAddress: int(address.ID),
			Recipient:   address.Recipient,
			PhoneNumber: address.PhoneNumber,
			Address:     address.Address,
			Note:        address.Note,
			Mark:        address.Mark,
			IsPrimary:   address.IsPrimary,
			UserId:      int(address.UserId),
		}

		addressResponses = append(addressResponses, addressResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Alamat berhasil didapatkan",
		"Data":    addressResponses,
	})
}

func (ph *ProfileHandler) UpdateAddressProfile(c echo.Context) error {
	var address ut.UserAddress

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	address.UserId = 1

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

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": message,
			})
		}
	}

	checkPhoneNumber := ""
	for i := 0; i < len(address.PhoneNumber); i++ {
		if i == 2 {
			break
		}
		checkPhoneNumber += string(address.PhoneNumber[i])
	}

	if checkPhoneNumber != "08" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Pastikan nomor kamu dimulai dengan '08'",
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
		"Message": "Alamat berhasil diubah",
	})
}

func (ph *ProfileHandler) UpdatePasswordProfile(c echo.Context) error {
	var user ut.User
	var userPassword ut.UserPasswordRequest

	// var claims = midjwt.GetClaims2(c)
	// var userId = claims["user_id"].(float64)
	idUserSementara := 1

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
		})
	}

	if userPassword.Password != userPassword.ConfirmNewPassword {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Password tidak cocok",
		})
	}

	message, err := ph.profileUsecase.UpdatePasswordProfile(&user, userPassword.OldPassword, userPassword.Password, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": message,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Password berhasil diubah",
	})
}

func (ph *ProfileHandler) LogoutProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
