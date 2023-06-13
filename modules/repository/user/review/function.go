package review

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
)

func (rr *reviewRepo) CreateReview(review *er.Review) error {
	if err := rr.db.Create(&review).Error; err != nil {
		return err
	}

	return nil
}
