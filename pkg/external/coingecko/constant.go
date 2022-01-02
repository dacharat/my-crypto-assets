package coingecko

type Chain string
type CoingeckoMapper map[string]string

const (
	Algo Chain = "algo"
)

var (
	AlgoCoinID = CoingeckoMapper{
		"PLANET": "planetwatch",
		"ALGO":   "algorand",
	}

	chain = map[Chain]CoingeckoMapper{
		Algo: AlgoCoinID,
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
