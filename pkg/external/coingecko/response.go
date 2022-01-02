package coingecko

type GetPriceResponse map[string]Price

type Price struct {
	USD float64 `json:"usd"`
}

func (p GetPriceResponse) Price(id string) float64 {
	return p[id].USD
}
