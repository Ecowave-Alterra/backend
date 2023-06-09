package ecommerce

type ProductResponse struct {
	Name            string
	Price           float64
	Rating          float64
	ProductImageUrl string
}

type ProductDetailResponse struct {
	Name            string
	Category        string
	Stock           uint
	Price           float64
	Status          string
	Rating          float64
	Description     string
	ProductImageUrl []string
}
