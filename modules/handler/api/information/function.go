package information

import (
	"math"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (informationHandler *InformationHandler) GetAllInformations() echo.HandlerFunc {
	return func(e echo.Context) error {
		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		informations, total, err := informationHandler.informationUsecase.GetAllInformations(offset, pageSize)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if page > int(math.Ceil(float64(total)/float64(pageSize))) {
			return e.JSON(http.StatusNotFound, echo.Map{
				"Message": "Not Found",
			})
		}

		if informations == nil || len(*informations) == 0 {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Belum ada list informasi",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informations,
			"Page":         page,
			"ToalPage":     int(math.Ceil(float64(total) / float64(pageSize))),
		})
	}
}

func (informationHandler *InformationHandler) GetInformationById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "input id is not a number",
			})
		}

		information, err = informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Information": information,
		})
	}
}

func (informationHandler *InformationHandler) CreateInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		statusIDStr := e.FormValue("StatusId")
		statusID, err := strconv.ParseUint(statusIDStr, 10, 64)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid StatusId",
			})
		}

		fileHeader, err := e.FormFile("PhotoContentUrl")
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Mohon maaf, Anda harus mengungga foto",
			})
		}

		fileExtension := filepath.Ext(fileHeader.Filename)
		allowedExtensions := map[string]bool{
			".png":  true,
			".jpeg": true,
			".jpg":  true,
		}
		if !allowedExtensions[fileExtension] {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Mohon maaf, format file yang Anda unggah tidak sesuai",
			})
		}

		maxFileSize := 4 * 1024 * 1024
		fileSize := fileHeader.Size
		if fileSize > int64(maxFileSize) {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB",
			})
		}

		PhotoUrl, _ := cloudstorage.UploadToBucket(e.Request().Context(), fileHeader)

		information := &ei.Information{
			Title:           e.FormValue("Title"),
			Content:         e.FormValue("Content"),
			PhotoContentUrl: PhotoUrl,
			StatusId:        uint(statusID),
		}

		if err := e.Validate(information); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "max" && e.Field() == "Title" {
						message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
					}
				}
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": message,
				})
			}
		}

		err = informationHandler.informationUsecase.CreateInformation(information)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if information.StatusId == 2 {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil menambahkan informasi ke dalam draft",
			})
		} else {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil menerbitkan informasi baru",
			})
		}
	}
}

func (informationHandler *InformationHandler) UpdateInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "input id is not a number",
			})
		}

		informationBefore, err := informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		information, err = informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		if err := e.Bind(&information); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid Request Body",
			})
		}

		if err := e.Validate(information); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "max" && e.Field() == "Title" {
						message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
					}
				}
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": message,
				})
			}
		}

		err = informationHandler.informationUsecase.UpdateInformation(int(information.ID), information)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
			})
		}

		if information.StatusId == 2 {
			if informationBefore.StatusId != information.StatusId {
				return e.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Informasi berhasil dipindahkan ke dalam draft",
				})
			}
		} else if information.StatusId == 1 {
			if informationBefore.StatusId != information.StatusId {
				return e.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Anda berhasil menerbitkan informasi baru",
				})
			}
		}
		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Anda berhasil mengubah informasi",
		})
	}
}

func (informationHandler *InformationHandler) DeleteInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ei.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Messages": "input id is not a number",
			})
		}

		information, err = informationHandler.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		err = informationHandler.informationUsecase.DeleteInformation(int(information.InformationId))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Anda berhasil menghapus informasi",
		})
	}
}

func (informationHandler *InformationHandler) SearchInformations() echo.HandlerFunc {
	return func(e echo.Context) error {
		var informations *[]ei.Information
		var err error

		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		keyword := e.QueryParam("keyword")
		informations, total, err := informationHandler.informationUsecase.SearchInformations(keyword, offset, pageSize)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if len(*informations) == 0 {
			return e.JSON(http.StatusOK, echo.Map{
				"Message": "Informasi yang anda cari tidak ditemukan",
			})
		} else {
			if page > int(math.Ceil(float64(total)/float64(pageSize))) {
				return e.JSON(http.StatusNotFound, echo.Map{
					"Message": "Not Found",
				})
			}

			return e.JSON(http.StatusOK, map[string]interface{}{
				"Informations": informations,
				"Page":         page,
				"TotalPage":    int(math.Ceil(float64(total) / float64(pageSize))),
			})
		}
	}
}

func (informationHandler *InformationHandler) FilterInformations() echo.HandlerFunc {
	return func(e echo.Context) error {
		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		keywordString := e.QueryParam("keyword")
		keyword, _ := strconv.Atoi(keywordString)
		informations, total, err := informationHandler.informationUsecase.FilterInformations(keyword, offset, pageSize)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if keyword == 1 && len(*informations) == 0 {
			return e.JSON(http.StatusOK, echo.Map{
				"Message": "Belum ada informasi yang terbit",
			})
		} else if keyword == 2 && len(*informations) == 0 {
			return e.JSON(http.StatusOK, echo.Map{
				"Message": "Belum ada informasi dalam draft",
			})
		} else if keyword == 1 || keyword == 2 {
			if page > int(math.Ceil(float64(total)/float64(pageSize))) {
				return e.JSON(http.StatusNotFound, echo.Map{
					"Message": "Not Found",
				})
			}

			return e.JSON(http.StatusOK, map[string]interface{}{
				"Informations": informations,
				"Page":         page,
				"TotalPage":    int(math.Ceil(float64(total) / float64(pageSize))),
			})
		}

		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid parameters",
		})
	}
}
