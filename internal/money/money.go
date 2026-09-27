package money

import (
	"fmt"
)

var _ fmt.Stringer = Money{}

var currencies = map[string]struct {
	Divisor  int64
	Decimals int
}{
	"EUR": {100, 2},
	"USD": {100, 2},
	"GBP": {100, 2},
	"JPY": {1, 0},
}

type Money struct {
	Minor    int64  `json:"minor"`
	Currency string `json:"currency"`
}

func (m Money) FormatChecked() (string, error) {

	// Known currency?
	d, ok := currencies[m.Currency]

	if !ok {
		return "", fmt.Errorf("currency %s unknown not found in decimals", m.Currency)
	}

	sign := ""
	if m.Minor < 0 {
		sign = "-"
		m.Minor = -m.Minor
	}

	if d.Decimals == 0 {
		return fmt.Sprintf("%s %s%d", m.Currency, sign, m.Minor), nil
	}

	preDecimal := (m.Minor / d.Divisor)
	postDecimal := (m.Minor % d.Divisor)

	return fmt.Sprintf("%s %s%d.%0*d", m.Currency, sign, preDecimal, d.Decimals, postDecimal), nil
}

func (m Money) String() string {
	out, err := m.FormatChecked()

	if err != nil {
		return fmt.Sprintf("INVALID(%s %d) error: %s", m.Currency, m.Minor, err)
	}

	return out
}
