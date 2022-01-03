package binanceservice

type Account struct {
	Assets     Assets  `json:"assets"`
	TotalPrice float64 `json:"totalPrice"`
}

type Assets []*Asset

type Asset struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
}

func (a Assets) TotalPrice() float64 {
	var total float64

	for _, asset := range a {
		total += asset.Price
	}

	return total
}
