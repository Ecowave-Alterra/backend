package review

import (
	"fmt"
	"log"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) CreateReview(c echo.Context) error {
	cloudstorage.Folder = "img/reviews/"
	cloudstorage.FolderVideo = "video/reviews/"

	transactionId := c.Param("id")

	ratingService, err := strconv.ParseFloat(c.FormValue("RatingService"), 64)
	if err != nil {
		return err
	}

	if err := rh.reviewUsecase.CreateReview(ratingService, transactionId); err != nil {
		return c.JSON(500, echo.Map{
			"Message": err,
		})
	}

	countTransactionDetail, err := rh.reviewUsecase.CountTransactionDetail(transactionId)
	if err != nil {
		return c.JSON(500, echo.Map{
			"Message": err,
		})
	}

	log.Println(countTransactionDetail)

	var ratingProduct float64
	for i := 1; i <= countTransactionDetail; i++ {
		ratingProduct, err = strconv.ParseFloat(c.FormValue(fmt.Sprintf("RatingProduct%d", i)), 64)
		if err != nil {
			return err
		}
		comment := c.FormValue(fmt.Sprintf("Comment%d", i))
		fileHeader, _ := c.FormFile(fmt.Sprintf("PhotoUrl%d", i))
		videoHeader, _ := c.FormFile(fmt.Sprintf("VideoUrl%d", i))

		if err := rh.reviewUsecase.CreateReviewDetail(ratingProduct, comment, fileHeader, videoHeader, transactionId); err != nil {
			return c.JSON(500, echo.Map{
				"Message": err,
			})
		}
	}

	idUserTemp := 2
	if ratingService != 0 && ratingProduct != 0 {
		if err := rh.reviewUsecase.UpdatePoint(idUserTemp); err != nil {
			return err
		}
	}

	return c.JSON(200, echo.Map{
		"Message": "Success create review",
	})
}
