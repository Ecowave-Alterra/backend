package review

import (
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
)

func (rc *reviewUsecase) GetAllProducts(offset, pageSize int) ([]re.GetAllReviewResponse, int64, error) {
	return rc.reviewRepo.GetAllProducts(offset, pageSize)
}

func (rc *reviewUsecase) SearchProduct(search string, offset, pageSize int) ([]re.GetAllReviewResponse, int64, error) {
	return rc.reviewRepo.SearchProduct(search, offset, pageSize)
}

func (rc *reviewUsecase) GetProductReviewById(productId string, offset, pageSize int) ([]re.ReviewResponse, int64, error) {
	return rc.reviewRepo.GetProductReviewById(productId, offset, pageSize)
}
