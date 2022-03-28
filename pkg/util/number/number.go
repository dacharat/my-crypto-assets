package number

import (
	"math"
	"math/big"
)

func ToFloat(amount, decimal int) float64 {
	return float64(amount) / math.Pow(10, float64(decimal))
}

func BigIntToFloat(amount *big.Int, decimal int) float64 {
	bFloat := new(big.Float).SetInt(amount)
	f64, _ := new(big.Float).Quo(bFloat, big.NewFloat(math.Pow(10, float64(decimal)))).Float64()

	return f64
}

func LargeStringToFloat(amount string, decimal int) float64 {
	b := new(big.Int)
	bFloat, ok := b.SetString(amount, 10)
	if !ok {
		return 0
	}

	return BigIntToFloat(bFloat, decimal)
}
