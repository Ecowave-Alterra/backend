package order

type OrderResponse struct {
	ProductImageUrl      string
	PaymentStatus        string
	ProductName          string
	ProductQty           uint
	ProductPrice         float64
	TransactionDetailQty uint
	Total                float64
}

type OrderDetailResponse struct {
	TransactionId   uint
	ProductId       uint
	Qty             uint
	SubTotalPrice   float64
	NameProduct     string
	ProductImageUrl string
}

type Order struct {
	UserId             uint
	PaymentMethod      string
	ExpeditionId       uint
	VoucherId          uint
	PromoName          string
	AddressId          uint
	StatusTransaction  string
	ShippingCost       float64
	ProductCost        float64
	Point              float64
	TotalPrice         float64
	TotalProduct       uint
	TransactionDetails []OrderDetailResponse
}
