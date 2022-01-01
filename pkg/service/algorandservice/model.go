package algorandservice

type Account struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Amount float64 `json:"amount"`
	ID     int     `json:"id"`
	Name   string  `json:"name"`
}
