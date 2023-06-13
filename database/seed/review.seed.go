package seed

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
)

func CreateReview() []*er.Review {
	reviews := []*er.Review{
		{
			Rating:       4.5,
			Comment:      "Bagus",
			CommentAdmin: "Terimakasih reviewnya",
			PhotoUrl:     "https://storage.googleapis.com/ecowave/img/review.png",
			VideoUrl:     "https://storage.googleapis.com/ecowave/img/review.mp4",
		},
		{
			Rating:       2.5,
			Comment:      "Jelek",
			CommentAdmin: "Terimakasih reviewnya",
			PhotoUrl:     "https://storage.googleapis.com/ecowave/img/review.png",
			VideoUrl:     "https://storage.googleapis.com/ecowave/img/review.mp4",
		},
		{
			Rating:       5,
			Comment:      "Sangat bagus!!!",
			CommentAdmin: "Terimakasih reviewnya",
			PhotoUrl:     "https://storage.googleapis.com/ecowave/img/review.png",
			VideoUrl:     "https://storage.googleapis.com/ecowave/img/review.mp4",
		},
		{
			Rating:       4,
			Comment:      "Bahannya berkualitas",
			CommentAdmin: "Terimakasih reviewnya",
			PhotoUrl:     "https://storage.googleapis.com/ecowave/img/review.png",
			VideoUrl:     "https://storage.googleapis.com/ecowave/img/review.mp4",
		},
	}

	return reviews
}
