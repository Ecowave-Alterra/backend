package seed

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
)

func CreateReview() []*er.RatingProduct {
	ratingProducts := []*er.RatingProduct{
		{
			Rating:              4.5,
			Comment:             "Bagus",
			CommentAdmin:        "Terimakasih Reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 1,
		},
		{
			Rating:              3,
			Comment:             "Lumayan",
			CommentAdmin:        "Terimakasih Reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 2,
		},
		{
			Rating:              2.5,
			Comment:             "Jelek",
			CommentAdmin:        "Terimakasih Reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 3,
		},
		{
			Rating:              4.5,
			Comment:             "Perfect",
			CommentAdmin:        "Terimakasih Reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 4,
		},
	}

	return ratingProducts
}
