package shared

type Account struct {
	Address    string  `json:"address,omitempty"`
	Assets     Assets  `json:"assets"`
	TotalPrice float64 `json:"totalPrice"`
}

type Assets []*Asset
type Asset struct {
	ID            int     `json:"id,omitempty"`
	Amount        float64 `json:"amount"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	FormatedPrice string  `json:"formatedPrice,omitempty"`
}

func (a Assets) TotalPrice() float64 {
	var total float64

	for _, asset := range a {
		total += asset.Price
	}

	return total
}
