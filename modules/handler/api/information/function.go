package information

import (
	"encoding/csv"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (ih *InformationHandler) GetAllInformations() echo.HandlerFunc {
	return func(e echo.Context) error {
		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		informations, total, err := ih.informationUsecase.GetAllInformations(offset, pageSize)
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

func (ih *InformationHandler) GetInformationById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ie.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "input id is not a number",
			})
		}

		information, err = ih.informationUsecase.GetInformationById(id)
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

func (ih *InformationHandler) CreateInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		cloudstorage.Folder = "img/informations/"

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

		information := &ie.Information{
			Title:           e.FormValue("Title"),
			Content:         e.FormValue("Content"),
			PhotoContentUrl: PhotoUrl,
			Status:          e.FormValue("Status"),
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

		err = ih.informationUsecase.CreateInformation(information)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		if strings.EqualFold(information.Status, "Draft") {
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

func (ih *InformationHandler) UpdateInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		cloudstorage.Folder = "img/informations/"

		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "input id is not a number",
			})
		}

		informationBefore, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		information, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		title := e.FormValue("Title")
		content := e.FormValue("Content")
		status := e.FormValue("Status")
		fileHeader, err := e.FormFile("PhotoContentUrl")

		if title != "" {
			information.Title = title
		}
		if content != "" {
			information.Content = content
		}
		if status != "" {
			information.Status = status
		}
		if fileHeader != nil {
			if informationBefore.PhotoContentUrl != "" {
				fileName, _ := cloudstorage.GetFileName(informationBefore.PhotoContentUrl)
				if err != nil {
					return e.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Gagal mendapatkan nama file",
					})
				}
				err = cloudstorage.DeleteImage(fileName)
				if err != nil {
					return e.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Gagal menghapus file pada cloud storage",
					})
				}
			}

			PhotoUrl, _ := cloudstorage.UploadToBucket(e.Request().Context(), fileHeader)
			information.PhotoContentUrl = PhotoUrl
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

		err = ih.informationUsecase.UpdateInformation(int(information.InformationId), information)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
			})
		}

		if strings.EqualFold(information.Status, "Draft") {
			if informationBefore.Status != information.Status {
				return e.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Informasi berhasil dipindahkan ke dalam draft",
				})
			}
		} else if strings.EqualFold(information.Status, "Terbit") {
			if informationBefore.Status != information.Status {
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

func (ih *InformationHandler) DeleteInformation() echo.HandlerFunc {
	return func(e echo.Context) error {
		var information *ie.Information
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"Messages": "input id is not a number",
			})
		}

		information, err = ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Record Not Found",
			})
		}

		photoContentUrl := information.PhotoContentUrl
		if photoContentUrl != "" {
			fileName, err := cloudstorage.GetFileName(photoContentUrl)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal mendapatkan nama file",
				})
			}
			err = cloudstorage.DeleteImage(fileName)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal menghapus file pada cloud storage",
				})
			}
		}

		err = ih.informationUsecase.DeleteInformation(int(information.InformationId))
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

func (ih *InformationHandler) SearchInformations() echo.HandlerFunc {
	return func(e echo.Context) error {
		var informations *[]ie.Information
		var err error

		pageParam := e.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		search := e.QueryParam("search")
		filter := e.QueryParam("filter")

		validParams := map[string]bool{"search": true, "filter": true, "page": true}
		for param := range e.QueryParams() {
			if !validParams[param] {
				return e.JSON(http.StatusBadRequest, echo.Map{
					"Message": "Invalid parameter",
				})
			}
		}

		informations, total, err := ih.informationUsecase.SearchInformations(search, filter, offset, pageSize)
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

func (ih *InformationHandler) DownloadCSVFile() echo.HandlerFunc {
	return func(e echo.Context) error {
		informations, err := ih.informationUsecase.GetAllInformationsNoPagination()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
			})
		}

		file, err := os.Create("information-data.csv")
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to create CSV file",
				"Error":   err,
			})
		}

		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				log.Println("Error closing file:", closeErr)
			}
		}()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		csvHeader := []string{"InformationId", "Title", "Content", "Status", "ViewCount", "BookmarkCount", "PhotoContentUrl"}
		err = writer.Write(csvHeader)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to write CSV header",
				"Error":   err,
			})
		}

		for _, info := range *informations {
			record := []string{
				strconv.Itoa(int(info.InformationId)),
				info.Title,
				info.Content,
				info.Status,
				strconv.Itoa(int(info.ViewCount)),
				strconv.Itoa(int(info.BookmarkCount)),
				info.PhotoContentUrl,
			}

			err = writer.Write(record)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to write CSV record",
					"Error":   err,
				})
			}
		}

		writer.Flush()
		if err := writer.Error(); err != nil {
			return e.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to flush CSV writer",
				"Error":   err,
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Successfully generate CSV file",
		})
	}
}
