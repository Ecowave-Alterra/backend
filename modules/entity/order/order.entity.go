package order

type OrderResponse struct {
	ProductImageUrl string
	PaymentStatus   string
	ProductName     string
	ProductQty      uint
	ProductPrice    float64
	TotalQty        uint
	Total           float64
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
	ExpeditionName   string
	ReceiptNumber    string
	ExpeditionStatus string
	AddressId        uint
	ShippingCost     float64

	PromoName         string
	TotalProduct      uint
	ProductCost       float64
	PaymentMethod     string
	StatusTransaction string
	Point             float64
	TotalPrice        float64
	VoucherId         uint
	Discount          float64
	OrderDetails      []OrderDetailResponse
}
