package review

import (
	"context"
	"mime/multipart"

	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	// et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
)

func (rc *reviewUsecase) CreateReview(ratingService float64, transactionId string) error {
	idTransaction, err := rc.reviewRepo.GetIdTransaction(transactionId)
	if err != nil {
		return err
	}

	review := er.Review{
		RatingService: ratingService,
		TransactionId: uint(idTransaction),
	}

	return rc.reviewRepo.CreateReview(&review)
}

func (rc *reviewUsecase) CreateReviewDetail(ratingProduct float64, comment string, fileHeader, videoHeader *multipart.FileHeader, transactionId string) error {
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

	idTransaction, err := rc.reviewRepo.GetIdTransaction(transactionId)
	if err != nil {
		return err
	}

	idReview, err := rc.reviewRepo.GetIdReview(idTransaction)
	if err != nil {
		return err
	}

	reviewDetail := er.ReviewDetail{
		RatingProduct: ratingProduct,
		Comment:       comment,
		PhotoUrl:      photoUrl,
		VideoUrl:      videoUrl,
		ReviewId:      uint(idReview),
		// ReviewId: 7,
	}

	return rc.reviewRepo.CreateReviewDetail(&reviewDetail)
}

func (rc *reviewUsecase) CountTransactionDetail(transactionId string) (int, error) {
	return rc.reviewRepo.CountTransactionDetail(transactionId)
}

func (rc *reviewUsecase) UpdatePoint(id int) error {
	point, err := rc.reviewRepo.GetPoint(id)
	if err != nil {
		return err
	}

	point += 10

	if err := rc.reviewRepo.UpdatePoint(id, point); err != nil {
		return err
	}

	return nil
}
