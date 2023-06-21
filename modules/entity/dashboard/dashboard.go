package dashboard

type FavouriteProducts struct {
	Name        string
	TotalOrders int64
}

type TopReviews struct {
	Name         string
	TotalReviews int64
}
type ChartResponse struct {
	Label string
	Value float64
}
