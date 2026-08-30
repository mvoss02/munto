package money

import (
	"testing"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		name     string
		currency string
		in       int64
		want     string
	}{
		{"classic 4 digits", "EUR", 1234, "EUR 12.34"},
		{"cents", "USD", 5, "USD 0.05"},
		{"classic 3 digits", "EUR", 999, "EUR 9.99"},
		{"negative", "GBP", -7432, "GBP -74.32"},
		{"negative cents", "GBP", -5, "GBP -0.05"},
		{"zero", "USD", 0, "USD 0.00"},
		{"whole amounts", "EUR", 500, "EUR 5.00"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Format(tc.in, tc.currency)
			if got != tc.want {
				t.Errorf("Format(%d, %s) = %s, want %s", tc.in, tc.currency, got, tc.want)
			}
		})
	}
}
