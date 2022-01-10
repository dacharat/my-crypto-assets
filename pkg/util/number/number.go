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
