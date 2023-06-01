package profile

import (
	"log"
	"net/http"

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

		log.Println("aha")
		profilePhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)
		log.Println(profilePhotoUrl)
		userDetail.ProfilePhotoUrl = profilePhotoUrl
		log.Println("aha")
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

// func (ph *ProfileHandler) GetUserDetailProfile(c echo.Context) error {
// 	var userDetail ut.UserDetail

// 	if err := ph.profileUsecase.GetUserDetailProfile(&userDetail, 1); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "fail",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success",
// 		"data":    userDetail,
// 	})
// }
