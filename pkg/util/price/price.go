package price

import (
	"fmt"
	"strings"
)

func Dollar(amount float64) string {
	num := fmt.Sprintf("%.2f", amount)
	return fmt.Sprintf("$%s", format(num))
}

func format(number string) string {
	number = strings.ReplaceAll(number, ",", "")

	decimal := ""

	if strings.Index(number, ".") != -1 {
		decimal = number[strings.Index(number, ".")+1:]
		number = number[0:strings.Index(number, ".")]
	}

	for i := 0; i <= len(number); i += 4 {
		a := number[0 : len(number)-i]
		b := number[len(number)-i:]
		number = a + "," + b
	}

	if number[0:1] == "," {
		number = number[1:]
	}

	if number[len(number)-1:] == "," {
		number = number[0 : len(number)-1]
	}

	if decimal != "" {
		number = number + "." + decimal
	}

	return number
}
