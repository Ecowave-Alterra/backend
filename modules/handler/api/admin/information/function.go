package information

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if total != 0 {
			if page > totalPages {
				return c.JSON(http.StatusNotFound, echo.Map{
					"Message": "Halaman Tidak Ditemukan",
					"Status":  http.StatusNotFound,
				})
			}
		} else {
			page = 0
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Status":  404,
				"Message": "Belum ada list informasi",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":       200,
			"Page":         page,
			"TotalPage":    totalPages,
			"Informations": informations,
		})
	}
}

func (ih *InformationHandler) GetInformationById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var information *ie.Information
		id := c.Param("id")

		information, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
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

		fileHeader, _ := c.FormFile("PhotoContentUrl")
		title := c.FormValue("Title")
		content := c.FormValue("Content")
		status := c.FormValue("Status")

		if strings.EqualFold(status, "Draft") {
			err := ih.informationUsecase.CreateInformationDraft(fileHeader, title, content, status)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": err.Error(),
					"Status":  http.StatusInternalServerError,
				})
			}
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"Message": "Anda berhasil menambahkan informasi ke dalam draft",
				"Status":  http.StatusCreated,
			})
		} else {
			err := ih.informationUsecase.CreateInformation(fileHeader, title, content, status)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": err.Error(),
					"Status":  http.StatusInternalServerError,
				})
			}
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"Message": "Anda berhasil menerbitkan informasi baru",
				"Status":  http.StatusCreated,
			})
		}
	}
}

func (ih *InformationHandler) UpdateInformation() echo.HandlerFunc {
	return func(c echo.Context) error {
		cloudstorage.Folder = "img/informations/"

		id := c.Param("id")

		informationBefore, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		information, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		title := c.FormValue("Title")
		content := c.FormValue("Content")
		status := c.FormValue("Status")
		fileHeader, _ := c.FormFile("PhotoContentUrl")

		err = ih.informationUsecase.UpdateInformation(informationBefore, information, fileHeader, title, content, status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
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
		id := c.Param("id")

		information, err := ih.informationUsecase.GetInformationById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		photoContentURL := information.PhotoContentUrl
		if photoContentURL != "" {
			fileName := cloudstorage.GetFileName(photoContentURL)
			err = cloudstorage.DeleteImage(fileName)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": err.Error(),
					"Status":  http.StatusInternalServerError,
				})
			}
		}

		err = ih.informationUsecase.DeleteInformation(information.InformationId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
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
					"Message": "Masukkan parameter dengan benar",
					"Status":  http.StatusBadRequest,
				})
			}
		}

		informations, total, err := ih.informationUsecase.SearchInformations(search, filter, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		draftInformations := make([]ie.Information, 0)
		for _, info := range *informations {
			if strings.EqualFold(info.Status, "Draft") {
				draftInformations = append(draftInformations, info)
			}
		}

		terbittInformations := make([]ie.Information, 0)
		for _, info := range *informations {
			if strings.EqualFold(info.Status, "Terbit") {
				terbittInformations = append(terbittInformations, info)
			}
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if total != 0 {
			if page > totalPages {
				return c.JSON(http.StatusNotFound, echo.Map{
					"Message": "Halaman Tidak Ditemukan",
					"Status":  http.StatusNotFound,
				})
			}
		} else {
			page = 0
			if search == "" {
				if len(draftInformations) == 0 && strings.EqualFold(filter, "Draft") {
					return c.JSON(http.StatusNotFound, echo.Map{
						"Message": "Belum ada informasi dalam draft",
						"Status":  http.StatusNotFound,
					})
				} else if len(terbittInformations) == 0 && strings.EqualFold(filter, "Terbit") {
					return c.JSON(http.StatusNotFound, echo.Map{
						"Message": "Belum ada informasi yang terbit",
						"Status":  http.StatusNotFound,
					})
				}
			} else {
				return c.JSON(http.StatusNotFound, echo.Map{
					"Message": "Informasi yang anda cari tidak ditemukan",
					"Status":  http.StatusNotFound,
				})
			}
		}

		if len(*informations) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Status":  404,
				"Message": "Belum ada list informasi",
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

func (ih *InformationHandler) DownloadCSVFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		informations, err := ih.informationUsecase.GetAllInformationsNoPagination()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil membuat file CSV",
			"Status":  http.StatusOK,
			"Data":    informations,
		})
	}
}
