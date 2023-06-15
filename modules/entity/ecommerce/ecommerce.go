package ecommerce

type ProductResponse struct {
	Name            string
	Price           float64
	Rating          float64
	ProductImageUrl string
}

type QueryResponse struct {
	Id              int     `json:"Id"`
	ProductId       string  `json:"ProductId"`
	Name            string  `json:"Name"`
	Category        string  `json:"Category"`
	Stock           int     `json:"Stock"`
	Price           float64 `json:"Price"`
	Status          string  `json:"Status"`
	Description     string  `json:"Description"`
	FullName        string  `json:"Full_name"`
	ProfilePhotoUrl string  `json:"ProfilePhotoUrl"`
	Rating          float64 `json:"Rating"`
	Comment         string  `json:"Comment"`
	CommentAdmin    string  `json:"Comment_admin"`
	PhotoURL        string  `json:"Photo_url"`
	VideoURL        string  `json:"Video_url"`
}

type ReviewResponse struct {
	FullName        string  `json:"FullName"`
	ProfilePhotoUrl string  `json:"ProfilePhotoUrl"`
	Rating          float64 `json:"Rating"`
	Comment         string  `json:"Comment"`
	CommentAdmin    string  `json:"Comment_admin"`
	PhotoURL        string  `json:"Photo_url"`
	VideoURL        string  `json:"Video_url"`
}

type ProductDetailResponse struct {
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
