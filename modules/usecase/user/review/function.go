package review

import (
	"context"
	"mime/multipart"

	er "github.com/berrylradianh/ecowave-go/modules/entity/review"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
)

func (rc *reviewUsecase) CountTransactionDetail(transactionId string) (int, error) {
	return rc.reviewRepo.CountTransactionDetail(transactionId)
}

func (rc *reviewUsecase) GetIdTransactionDetail(transactionId string) ([]int, error) {
	idTransaction, err := rc.reviewRepo.GetIdTransaction(transactionId)
	if err != nil {
		return nil, err
	}

	productIds, err := rc.reviewRepo.GetProductId(transactionId)
	if err != nil {
		return nil, err
	}

	var idTransactionDetails []int
	for _, productId := range productIds {
		idTransactionDetail, err := rc.reviewRepo.GetIdTransactionDetail(idTransaction, productId)
		if err != nil {
			return nil, err
		}

		idTransactionDetails = append(idTransactionDetails, idTransactionDetail)
	}

	return idTransactionDetails, nil
}

func (rc *reviewUsecase) CreateRatingProduct(rating float64, comment string, fileHeader, videoHeader *multipart.FileHeader, transactionDetailId int) error {
	var photoUrl string
	var videoUrl string
	var err error

	if fileHeader != nil {
		if err := vld.ValidateFileExtension(fileHeader); err != nil {
			return err
		}

		maxFileSize := 4 * 1024 * 1024
		if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
			return err
		}

		photoUrl, err = cloudstorage.UploadToBucket(context.Background(), fileHeader)
		if err != nil {
			return err
		}
	}

	if videoHeader != nil {
		if err := vld.ValidateVideoExtension(videoHeader); err != nil {
			return err
		}

		maxVideoSize := 4 * 1024 * 1024
		if err := vld.ValidateVideoSize(videoHeader, int64(maxVideoSize)); err != nil {
			return err
		}

		videoUrl, err = cloudstorage.UploadVideoToBucket(context.Background(), videoHeader)
		if err != nil {
			return err
		}
	}

	ratingProduct := er.RatingProduct{
		Rating:              rating,
		Comment:             comment,
		PhotoUrl:            photoUrl,
		VideoUrl:            videoUrl,
		TransactionDetailId: uint(transactionDetailId),
	}

	return rc.reviewRepo.CreateRatingProduct(&ratingProduct)
}

func (rc *reviewUsecase) UpdateExpeditionRating(ratingExpedition float32, transactionId string) error {
	return rc.reviewRepo.UpdateExpeditionRating(ratingExpedition, transactionId)
}

func (rc *reviewUsecase) UpdatePoint(idUser int) error {
	point, err := rc.reviewRepo.GetPoint(idUser)
	if err != nil {
		return err
	}

	point += 10

	if err := rc.reviewRepo.UpdatePoint(idUser, point); err != nil {
		return err
	}

	return nil
}
