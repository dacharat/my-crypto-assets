package price

import (
	"fmt"
)

func Dollar(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}
