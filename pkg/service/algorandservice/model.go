package algorandservice

type Account struct {
	Address    string `json:"address"`
	Assets     Assets `json:"assets"`
	TotalPirce string `json:"totalPrice"`
}

type Assets []Asset
type Asset struct {
	Amount        float64 `json:"amount"`
	ID            int     `json:"id,omitempty"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	FormatedPrice string  `json:"formatedPrice"`
}

func (a Assets) TotalPrice() float64 {
	var total float64

	for _, asset := range a {
		total += asset.Price
	}

	return total
}
