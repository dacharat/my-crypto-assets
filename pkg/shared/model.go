package shared

import "sort"

const (
	Algorand    Platform = "Algorand"
	Binance     Platform = "Binance"
	Bitkub      Platform = "Bitkub"
	BitkubChain Platform = "BitkubChain"
	BSC         Platform = "BinanceSmartChain"
	ElrondChain Platform = "ElrondChain"
)

var AvailablePlatform = []Platform{Algorand, Binance, Bitkub, BitkubChain, BSC, ElrondChain}

type Platform string

type Account struct {
	Platform   Platform `json:"platform"`
	Address    string   `json:"address,omitempty"`
	Assets     Assets   `json:"assets"`
	TotalPrice float64  `json:"totalPrice"`

	NeedCgkPrice bool `json:"-"`
}

type Assets []*Asset
type Asset struct {
	ID            int     `json:"id,omitempty"`
	Amount        float64 `json:"amount"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	FormatedPrice string  `json:"formatedPrice,omitempty"`
}

type GetAccountReq struct {
	WalletAddress string
}

func (a Assets) TotalPrice() float64 {
	var total float64

	for _, asset := range a {
		total += asset.Price
	}

	return total
}

func (a Assets) Sort() Assets {
	sort.Slice(a, func(i, j int) bool {
		return a[i].Price > a[j].Price
	})

	return a
}

func InAvailablePlatform(platfrom Platform) bool {
	for _, p := range AvailablePlatform {
		if p == platfrom {
			return true
		}
	}

	return false
}
