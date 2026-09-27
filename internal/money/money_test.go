package money

import (
	"strings"
	"testing"
)

func TestFormatChecked(t *testing.T) {
	cases_v1 := []struct {
		name    string
		money   Money
		want    string
		wantErr bool
	}{
		{"classic 4 digits", Money{Minor: 1234, Currency: "EUR"}, "EUR 12.34", false},
		{"cents", Money{Minor: 5, Currency: "USD"}, "USD 0.05", false},
		{"classic 3 digits", Money{Minor: 999, Currency: "EUR"}, "EUR 9.99", false},
		{"negative", Money{Minor: -7432, Currency: "GBP"}, "GBP -74.32", false},
		{"negative cents", Money{Minor: -5, Currency: "GBP"}, "GBP -0.05", false},
		{"zero", Money{Minor: 0, Currency: "USD"}, "USD 0.00", false},
		{"whole amounts", Money{Minor: 500, Currency: "EUR"}, "EUR 5.00", false},
		{"forced error", Money{Minor: 401, Currency: "AUD"}, "AUD 4.01", true},
		{"forced error", Money{Minor: 401, Currency: "JPY"}, "JPY 401", false},
	}

	for _, tc := range cases_v1 {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.money.FormatChecked()

			if err != nil && !tc.wantErr {
				t.Fatalf("got unwanted error: %s", err)
			}
			if err == nil && tc.wantErr {
				t.Errorf("expected error, got none")
			}

			if got != tc.want && !tc.wantErr {
				t.Errorf("Money{Minor: %d, Currency: %s} = %s, want %s", tc.money.Minor, tc.money.Currency, got, tc.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	cases_v2 := []struct {
		name    string
		money   Money
		want    string
		wantErr bool
	}{
		{"forced error", Money{Minor: 401, Currency: "AUD"}, "AUD 4.01", true},
		{"classic 4 digits", Money{Minor: 1234, Currency: "EUR"}, "EUR 12.34", false},
	}

	for _, tc := range cases_v2 {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.money.String()

			if tc.wantErr && !strings.Contains(got, "INVALID(") {
				t.Fatalf("error not caught got: %s", got)
			}

			if got != tc.want && !tc.wantErr {
				t.Errorf("Money{Minor: %d, Currency: %s} = %s, want %s", tc.money.Minor, tc.money.Currency, got, tc.want)
			}
		})
	}
}

func TestPassByValue(t *testing.T) {
	case_v3 := struct {
		name    string
		money   Money
		want    string
		wantErr bool
	}{"go vs python: passing struct/class", Money{Minor: -500, Currency: "EUR"}, "EUR -5.00", false}

	t.Run(case_v3.name, func(t *testing.T) {
		got := case_v3.money.String()

		if !case_v3.wantErr && got != case_v3.want {
			t.Errorf("Money{Minor: %d, Currency: %s} = %s, want %s", case_v3.money.Minor, case_v3.money.Currency, got, case_v3.want)
		}

		b := case_v3.money
		b.Minor = 0

		got_v2 := case_v3.money.String()

		if !case_v3.wantErr && got_v2 != case_v3.want {
			t.Errorf("Money{Minor: %d, Currency: %s} = %s, want %s", case_v3.money.Minor, case_v3.money.Currency, got_v2, case_v3.want)
		}
	})
}
