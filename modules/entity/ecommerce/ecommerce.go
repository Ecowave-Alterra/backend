package ecommerce

type ProductResponse struct {
	Name            string
	Price           float64
	Rating          float64
	ProductImageUrl string
}

type QueryResponse struct {
	Id           int     `json:"Id"`
	Name         string  `json:"Name"`
	Category     string  `json:"Category"`
	Stock        int     `json:"Stock"`
	Price        float64 `json:"Price"`
	Status       string  `json:"Status"`
	Description  string  `json:"Description"`
	FullName     string  `json:"Full_name"`
	Rating       int     `json:"Rating"`
	Comment      string  `json:"Comment"`
	CommentAdmin string  `json:"Comment_admin"`
	PhotoURL     string  `json:"Photo_url"`
	VideoURL     string  `json:"Video_url"`
}

type ReviewResponse struct {
	FullName     string  `json:"Full_name"`
	Rating       float32 `json:"Rating"`
	Comment      string  `json:"Comment"`
	CommentAdmin string  `json:"Comment_admin"`
	PhotoURL     string  `json:"Photo_url"`
	VideoURL     string  `json:"Video_url"`
}

type ProductDetailResponse struct {
	Name            string           `json:"Name"`
	Category        string           `json:"Category"`
	Stock           int              `json:"Stock"`
	Price           float64          `json:"Price"`
	Status          string           `json:"Status"`
	Description     string           `json:"Description"`
	ProductImageUrl []string         `json:"ProductImageUrl"`
	Review          []ReviewResponse `json:"Review"`
}
