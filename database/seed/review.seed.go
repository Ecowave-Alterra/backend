package seed

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
)

func CreateReview() []*er.RatingProduct {
	reviews := []*er.RatingProduct{
		{
			Rating:              4.5,
			Comment:             "Bagus",
			CommentAdmin:        "Terimakasih reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/img/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 1,
		},
		{
			Rating:              2.5,
			Comment:             "Jelek",
			CommentAdmin:        "Terimakasih reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/img/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 2,
		},
		{
			Rating:              5,
			Comment:             "Sangat bagus!!!",
			CommentAdmin:        "Terimakasih reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/img/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 3,
		},
		{
			Rating:              4,
			Comment:             "Bahannya berkualitas",
			CommentAdmin:        "Terimakasih reviewnya",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/img/reviews/review.png",
			VideoUrl:            "https://storage.googleapis.com/ecowave/video/reviews/review.mp4",
			TransactionDetailId: 4,
		},
	}

	return reviews
}
