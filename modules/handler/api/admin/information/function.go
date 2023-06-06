package information

import (
	"database/sql"
	"encoding/csv"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	"github.com/labstack/echo/v4"
)

func (ih *InformationHandler) GetAllInformations() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		informations, total, err := ih.informationUsecase.GetAllInformations(offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan informasi",
				"Status":  http.StatusInternalServerError,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		if informations == nil || len(*informations) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Belum ada list informasi",
				"Status":  http.StatusOK,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Informations": informations,
			"Page":         page,
			"TotalPage":    totalPages,
			"Status":       http.StatusOK,
		})
	}
}

func (ih *InformationHandler) GetInformationById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var information *ie.Information
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "ID harus berupa angka",
				"Status":  http.StatusBadRequest,
			})
		}

		information, err = ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan informasi",
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Information": information,
			"Status":      http.StatusOK,
		})
	}
}

func (ih *InformationHandler) CreateInformation() echo.HandlerFunc {
	return func(c echo.Context) error {
		cloudstorage.Folder = "img/informations/"

		fileHeader, err := c.FormFile("PhotoContentUrl")
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"Message": "Mohon maaf, Anda harus mengungga foto",
				"Status":  http.StatusUnprocessableEntity,
			})
		}

		if err := vld.ValidateFileExtension(fileHeader); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusUnsupportedMediaType,
			})
		}

		maxFileSize := 4 * 1024 * 1024
		if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
			return c.JSON(http.StatusRequestEntityTooLarge, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusRequestEntityTooLarge,
			})
		}

		PhotoUrl, err := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal upload image",
				"Status":  http.StatusInternalServerError,
			})
		}

		information := &ie.Information{
			Title:           c.FormValue("Title"),
			Content:         c.FormValue("Content"),
			PhotoContentUrl: PhotoUrl,
			Status:          c.FormValue("Status"),
		}

		err = ih.informationUsecase.CreateInformation(information)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		if strings.EqualFold(information.Status, "Draft") {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil menambahkan informasi ke dalam draft",
				"Status":  http.StatusOK,
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Anda berhasil menerbitkan informasi baru",
				"Status":  http.StatusOK,
			})
		}
	}
}

func (ih *InformationHandler) UpdateInformation() echo.HandlerFunc {
	return func(c echo.Context) error {
		cloudstorage.Folder = "img/informations/"

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "ID harus berupa angka",
				"Status":  http.StatusBadRequest,
			})
		}

		informationBefore, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Informasi tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		information, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Informasi tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		title := c.FormValue("Title")
		content := c.FormValue("Content")
		status := c.FormValue("Status")
		fileHeader, err := c.FormFile("PhotoContentUrl")

		if err := vld.ValidateFileExtension(fileHeader); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusUnsupportedMediaType,
			})
		}

		maxFileSize := 4 * 1024 * 1024
		if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
			return c.JSON(http.StatusRequestEntityTooLarge, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusRequestEntityTooLarge,
			})
		}

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
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Gagal mendapatkan nama file",
						"Status":  http.StatusInternalServerError,
					})
				}
				err = cloudstorage.DeleteImage(fileName)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Gagal menghapus file pada cloud storage",
						"Status":  http.StatusInternalServerError,
					})
				}
			}

			PhotoUrl, err := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal upload image",
					"Status":  http.StatusInternalServerError,
				})
			}
			information.PhotoContentUrl = PhotoUrl
		}

		err = ih.informationUsecase.UpdateInformation(int(information.InformationId), information)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
				"Status":  http.StatusBadRequest,
			})
		}

		if strings.EqualFold(information.Status, "Draft") {
			if informationBefore.Status != information.Status {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Informasi berhasil dipindahkan ke dalam draft",
					"Status":  http.StatusOK,
				})
			}
		} else if strings.EqualFold(information.Status, "Terbit") {
			if informationBefore.Status != information.Status {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"Message": "Anda berhasil menerbitkan informasi baru",
					"Status":  http.StatusOK,
				})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Anda berhasil mengubah informasi",
			"Status":  http.StatusOK,
		})
	}
}

func (ih *InformationHandler) DeleteInformation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var information *ie.Information
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "ID harus berupa angka",
				"Status":  http.StatusBadRequest,
			})
		}

		information, err = ih.informationUsecase.GetInformationById(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, echo.Map{
					"Message": "Informasi tidak ditemukan",
					"Status":  http.StatusNotFound,
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan informasi",
				"Status":  http.StatusInternalServerError,
			})
		}

		photoContentURL := information.PhotoContentUrl
		if photoContentURL != "" {
			fileName, err := cloudstorage.GetFileName(photoContentURL)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal mendapatkan nama file",
					"Status":  http.StatusInternalServerError,
				})
			}
			err = cloudstorage.DeleteImage(fileName)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal menghapus file pada cloud storage",
					"Status":  http.StatusInternalServerError,
				})
			}
		}

		err = ih.informationUsecase.DeleteInformation(int(information.InformationId))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Anda berhasil menghapus informasi",
			"Status":  http.StatusOK,
		})
	}
}

func (ih *InformationHandler) SearchInformations() echo.HandlerFunc {
	return func(c echo.Context) error {
		var informations *[]ie.Information
		var err error

		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		search := c.QueryParam("search")
		filter := c.QueryParam("filter")

		validParams := map[string]bool{"search": true, "filter": true, "page": true}
		for param := range c.QueryParams() {
			if !validParams[param] {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"Message": "Masukkan paramter dengan benar",
					"Status":  http.StatusBadRequest,
				})
			}
		}

		informations, total, err := ih.informationUsecase.SearchInformations(search, filter, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan informasi",
				"Status":  http.StatusInternalServerError,
			})
		}

		if len(*informations) == 0 {
			return c.JSON(http.StatusOK, echo.Map{
				"Message": "Informasi yang anda cari tidak ditemukan",
				"Status":  http.StatusOK,
			})
		} else {
			if page > int(math.Ceil(float64(total)/float64(pageSize))) {
				return c.JSON(http.StatusNotFound, echo.Map{
					"Message": "Not Found",
					"Status":  http.StatusNotFound,
				})
			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"Informations": informations,
				"Page":         page,
				"TotalPage":    int(math.Ceil(float64(total) / float64(pageSize))),
				"Status":       http.StatusOK,
			})
		}
	}
}

func (ih *InformationHandler) DownloadCSVFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		informations, err := ih.informationUsecase.GetAllInformationsNoPagination()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan informasi",
				"Status":  http.StatusInternalServerError,
			})
		}

		file, err := os.Create("information-data.csv")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal membuat file csv",
				"Status":  http.StatusInternalServerError,
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
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal membaca file csv",
				"Status":  http.StatusInternalServerError,
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
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal membaca file csv",
					"Status":  http.StatusInternalServerError,
				})
			}
		}

		writer.Flush()
		if err := writer.Error(); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal flush file",
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil membuat file csv",
			"Status":  http.StatusOK,
		})
	}
}
