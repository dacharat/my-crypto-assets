package number

import "math"

func ToFloat(amount, decimal int) float64 {
	return float64(amount) / math.Pow(10, float64(decimal))
}
