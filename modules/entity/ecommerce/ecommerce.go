package ecommerce

type ReviewResponse struct {
	Name         string  `json:"Name"`
	ProfilePhoto string  `json:"ProfilePhoto"`
	Rating       float64 `json:"Rating"`
	Comment      string  `json:"Comment"`
	CommentAdmin string  `json:"Comment_admin"`
	PhotoURL     string  `json:"Photo_url"`
	VideoURL     string  `json:"Video_url"`
}

type ProductResponse struct {
	ProductId       string           `json:"ProductId"`
	Name            string           `json:"Name"`
	Category        string           `json:"Category"`
	Stock           int              `json:"Stock"`
	Price           float64          `json:"Price"`
	Status          string           `json:"Status"`
	Description     string           `json:"Description"`
	ProductImageUrl []string         `json:"ProductImageUrl"`
	AvgRating       float64          `json:"AverageRating"`
	Review          []ReviewResponse `json:"Review"`
}
