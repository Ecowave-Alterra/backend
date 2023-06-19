package review

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) CreateReview(c echo.Context) error {
	cloudstorage.Folder = "img/reviews/"
	cloudstorage.FolderVideo = "video/reviews/"

	transactionId := c.Param("id")

	countTransactionDetail, err := rh.reviewUsecase.CountTransactionDetail(transactionId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Gagal",
			"Status":  http.StatusInternalServerError,
		})
	}

	idTransactionDetails, err := rh.reviewUsecase.GetIdTransactionDetail(transactionId)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Gagal mendapatkan id transaction detail",
			"Status":  http.StatusNotFound,
		})
	}

	var ratingProductF float64
	for i := 1; i <= countTransactionDetail; i++ {
		ratingProduct := c.FormValue(fmt.Sprintf("RatingProduct%d", i))
		if ratingProduct == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Masukkan rating produk",
				"Status":  http.StatusBadRequest,
			})
		} else {
			ratingProductF, err = strconv.ParseFloat(ratingProduct, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"Message": "Gagal",
					"Status":  http.StatusBadRequest,
				})
			}

			comment := c.FormValue(fmt.Sprintf("Comment%d", i))
			fileHeader, _ := c.FormFile(fmt.Sprintf("PhotoUrl%d", i))
			videoHeader, _ := c.FormFile(fmt.Sprintf("VideoUrl%d", i))

			if err := rh.reviewUsecase.CreateRatingProduct(ratingProductF, comment, fileHeader, videoHeader, idTransactionDetails[i-1]); err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal membuat rating",
					"Status":  http.StatusInternalServerError,
				})
			}
		}
	}

	var ratingExpeditionF float64
	ratingExpedition := c.FormValue("ExpeditionRating")
	if ratingExpedition == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Message": "Masukkan rating ekspedisi",
			"Status":  http.StatusBadRequest,
		})
	} else {
		ratingExpeditionF, err = strconv.ParseFloat(ratingExpedition, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Gagal",
				"Status":  http.StatusBadRequest,
			})
		}

		if err := rh.reviewUsecase.UpdateExpeditionRating(float32(ratingExpeditionF), transactionId); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mengubah rating ekspedisi",
				"Status":  http.StatusInternalServerError,
			})
		}
	}

	idUserTemp := 2
	if err := rh.reviewUsecase.UpdatePoint(idUserTemp); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Gagal mengubah point user",
			"Status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"Message": "Yey! Kamu mendapatkan point +10",
		"Status":  http.StatusCreated,
	})
}
