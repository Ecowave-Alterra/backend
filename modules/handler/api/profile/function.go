package profile

import (
	"net/http"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (ph *ProfileHandler) GetUserProfile(c echo.Context) error {
	var userProfileResponse *ut.UserResponse
	var err error

	idUserSementara := 1

	userProfileResponse, err = ph.profileUsecase.GetUserProfile(idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user profile",
		"data":    userProfileResponse,
	})
}

func (ph *ProfileHandler) GetUser2Profile(c echo.Context) error {
	var user2ProfileResponse *ut.User2Response
	var err error

	idUserSementara := 1

	user2ProfileResponse, _, _, err = ph.profileUsecase.GetUser2Profile(idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user 2 profile",
		"data":    user2ProfileResponse,
	})
}

func (ph *ProfileHandler) UpdateUserProfile(c echo.Context) error {
	idUserSementara := 1

	_, _, userDetailBefore, err := ph.profileUsecase.GetUser2Profile(idUserSementara)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	_, user, userDetail, err := ph.profileUsecase.GetUser2Profile(idUserSementara)
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

	if err := ph.profileUsecase.UpdateUserProfile(user, idUserSementara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := ph.profileUsecase.UpdateUserDetailProfile(userDetail, idUserSementara); err != nil {
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
	if isPrimary != address.IsPrimary {
		address.IsPrimary = isPrimary
	}

	if err := ph.profileUsecase.UpdateAddressProfile(&address, int(address.UserId), idAddress); err != nil {
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
