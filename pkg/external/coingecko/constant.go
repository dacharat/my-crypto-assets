package coingecko

type Chain string
type CoingeckoMapper map[string]string

const (
	Algo   Chain = "algo"
	Bitkub Chain = "bitkub"
)

var (
	AlgoCoinID = CoingeckoMapper{
		"PLANET": "planetwatch",
		"ALGO":   "algorand",
	}

	BitkubCoinID = CoingeckoMapper{
		"KUB":  "bitkub-coin",
		"KKUB": "bitkub-coin",
		"KBTC": "bitcoin",
	}

	chain = map[Chain]CoingeckoMapper{
		Algo:   AlgoCoinID,
		Bitkub: BitkubCoinID,
	}
)

func (c CoingeckoMapper) IDs() []string {
	ids := make([]string, len(c))
	i := 0
	for _, v := range c {
		ids[i] = v
		i++
	}
	return ids
}
