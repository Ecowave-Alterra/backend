package seed

import er "github.com/berrylradianh/ecowave-go/modules/entity/review"

func CreateReview() []*er.RatingProduct {
	reviews := []*er.RatingProduct{
		{
			Rating:              5.0,
			Comment:             "Kualitas bagus. Harga terjangkau",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/img/products/bottle.png",
			VideoUrl:            "https://storage.cloud.google.com/ecowave/video/reviews/video_2023-06-13_11-52-24.mp4",
			TransactionDetailId: 1,
		},
		{
			Rating:              3.7,
			Comment:             "Barang agak berbeda dengan deskripsi dan foto",
			PhotoUrl:            "https://storage.googleapis.com/ecowave/img/products/bottle.png",
			VideoUrl:            "https://storage.cloud.google.com/ecowave/video/reviews/video_2023-06-13_11-52-24.mp4",
			TransactionDetailId: 2,
		},
	}

	return reviews
}
