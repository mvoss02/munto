package money

import "fmt"

func Format(minor int64, currency string) string {
	sign := ""
	if minor < 0 {
		sign = "-"
		minor = -minor
	}

	// Assuming every currency comes with two decimal points
	preDecimal := (minor / 100)
	postDecimal := (minor % 100)

	return fmt.Sprintf("%s %s%d.%02d", currency, sign, preDecimal, postDecimal)
}
