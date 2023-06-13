package review

import (
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) CreateReview(c echo.Context) error {
	cloudstorage.Folder = "img/reviews/"
	cloudstorage.FolderVideo = "video/reviews/"

	rating, err := strconv.ParseFloat(c.FormValue("Rating"), 64)
	if err != nil {
		return err
	}
	comment := c.FormValue("Comment")
	fileHeader, _ := c.FormFile("PhotoUrl")
	videoHeader, _ := c.FormFile("VideoUrl")

	review := er.Review{
		Rating:  rating,
		Comment: comment,
	}

	if err := rh.reviewUsecase.CreateReview(&review, fileHeader, videoHeader); err != nil {
		return c.JSON(500, echo.Map{
			"Message": err,
		})
	}

	return c.JSON(200, echo.Map{
		"Message": "Success create review",
	})
}
