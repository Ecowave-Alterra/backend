package review

import (
	"context"
	"mime/multipart"

	er "github.com/berrylradianh/ecowave-go/modules/entity/review"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
)

func (rc *reviewUsecase) CreateReview(review *er.Review, fileHeader, videoHeader *multipart.FileHeader) error {
	if err := vld.ValidateFileExtension(fileHeader); err != nil {
		return err
	}

	maxFileSize := 4 * 1024 * 1024
	if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
		return err
	}

	photoUrl, err := cloudstorage.UploadToBucket(context.Background(), fileHeader)
	if err != nil {
		return err
	}

	if err := vld.ValidateVideoExtension(videoHeader); err != nil {
		return err
	}

	maxVideoSize := 4 * 1024 * 1024
	if err := vld.ValidateVideoSize(videoHeader, int64(maxVideoSize)); err != nil {
		return err
	}

	videoUrl, err := cloudstorage.UploadVideoToBucket(context.Background(), videoHeader)
	if err != nil {
		return err
	}

	reviewInput := er.Review{
		Rating:   review.Rating,
		Comment:  review.Comment,
		PhotoUrl: photoUrl,
		VideoUrl: videoUrl,
	}

	return rc.reviewRepo.CreateReview(&reviewInput)
}
