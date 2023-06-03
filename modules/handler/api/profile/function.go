package profile

import (
	"net/http"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (ph *ProfileHandler) GetUserProfile(c echo.Context) error {
	var user ut.User
	var userDetail ut.UserDetail

	idUserSementara := 5

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	available, err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if !available {
		userDetail.FullName = ""
		userDetail.EcoPoint = 0
	}

	userResponse := ut.UserResponse{
		FullName:    userDetail.FullName,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		EcoPoint:    userDetail.EcoPoint,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user profile",
		"data":    userResponse,
	})
}

func (ph *ProfileHandler) GetUser2Profile(c echo.Context) error {
	var user ut.User
	var userDetail ut.UserDetail

	idUserSementara := 2

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	available, err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if !available {
		userDetail.FullName = ""
		userDetail.ProfilePhotoUrl = ""
	}

	user2Response := ut.User2Response{
		FullName:        userDetail.FullName,
		Username:        user.Username,
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user profile",
		"data":    user2Response,
	})
}

func (ph *ProfileHandler) UpdateUserProfile(c echo.Context) error {
	var user ut.User
	var userDetail ut.UserDetail
	var userDetailBefore ut.UserDetail

	idUserSementara := 5

	if err := ph.profileUsecase.GetUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	available, err := ph.profileUsecase.GetUserDetailProfile(&userDetail, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	availableBefore, err := ph.profileUsecase.GetUserDetailProfile(&userDetailBefore, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
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
		userDetail.ProfilePhotoUrl = profilePhotoUrl
	}

	if err := ph.profileUsecase.UpdateUserProfile(&user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if !available && !availableBefore {
		userDetail = ut.UserDetail{
			FullName:        userDetail.FullName,
			ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
			UserId:          uint(idUserSementara),
		}

		if err := ph.profileUsecase.CreateUserDetailProfile(&userDetail); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "fail",
			})
		}
	}

	if err := ph.profileUsecase.UpdateUserDetailProfile(&userDetail, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user detail profile",
	})
}

func (ph *ProfileHandler) CreateAddressProfile(c echo.Context) error {
	var address ut.UserAddress
	address.UserId = 1

	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := ph.profileUsecase.CreateAddressProfile(&address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new address",
	})
}

func (ph *ProfileHandler) GetAllAddressProfile(c echo.Context) error {
	var address []ut.UserAddress
	idUserSementara := 1

	if err := ph.profileUsecase.GetAllAddressProfile(&address, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all address",
		"data":    address,
	})
}

func (ph *ProfileHandler) UpdateAddressProfile(c echo.Context) error {
	var address ut.UserAddress
	address.UserId = 1

	idAddress, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := ph.profileUsecase.GetAddressByIdProfile(&address, int(address.UserId), idAddress); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	recipient := c.FormValue("Recipient")
	phoneNumber := c.FormValue("PhoneNumber")
	addressFV := c.FormValue("Address")
	note := c.FormValue("Note")
	mark := c.FormValue("Mark")
	isPrimary, err := strconv.ParseBool(c.FormValue("IsPrimary"))
	if err != nil {
		return err
	}

	if recipient != "" {
		address.Recipient = recipient
	}
	if phoneNumber != "" {
		address.PhoneNumber = phoneNumber
	}
	if addressFV != "" {
		address.Address = addressFV
	}
	if note != "" {
		address.Note = note
	}
	if mark != "" {
		address.Mark = mark
	}

	if isPrimary == true {
		if err := ph.profileUsecase.UpdateAddressPrimaryProfile(&address, int(address.UserId)); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "fail",
			})
		}
	}
	address.IsPrimary = isPrimary

	if err := ph.profileUsecase.UpdateAddressByIdProfile(&address, int(address.UserId), idAddress); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update address by id",
	})
}

func (ph *ProfileHandler) UpdatePasswordProfile(c echo.Context) error {
	var user ut.User
	idUserSementara := 1

	oldPassword := c.FormValue("OldPassword")
	newPassword := c.FormValue("Password")
	confirmNewPassword := c.FormValue("ConfirmNewPassword")

	if newPassword != confirmNewPassword {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "password tidak cocok",
		})
	}

	message, err := ph.profileUsecase.UpdatePasswordProfile(&user, oldPassword, newPassword, idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": message,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "password berhasil diubah",
	})
}
