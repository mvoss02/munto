package ledger

import (
	"encoding/json"
	"testing"

	"github.com/mvoss02/munto/internal/money"
)

func TestTransactionDecode(t *testing.T) {
	cases := []struct {
		name    string
		json    string
		want    Transaction
		wantErr bool
	}{
		{"every key present", `{"id": "123", "account_id": "acc-1", "amount": {"minor": 500, "currency": "EUR"}, "date": "2026-12-12", "description": "money well spent"}`, Transaction{Id: "123", AccountID: "acc-1", Amount: money.Money{Minor: 500, Currency: "EUR"}, Date: "2026-12-12", Description: "money well spent"}, false},
		{"unknown extra key", `{"id": "123", "account_id": "acc-1", "amount": {"minor": 500, "currency": "EUR"}, "date": "2026-12-12", "description": "money well spent", "document": "x"}`, Transaction{Id: "123", AccountID: "acc-1", Amount: money.Money{Minor: 500, Currency: "EUR"}, Date: "2026-12-12", Description: "money well spent"}, false},
		{"missing optional key", `{"id": "123", "account_id": "acc-1", "amount": {"minor": 500, "currency": "EUR"}, "date": "2026-12-12"}`, Transaction{Id: "123", AccountID: "acc-1", Amount: money.Money{Minor: 500, Currency: "EUR"}, Date: "2026-12-12"}, false},
		{"wrong type", `{"id": "123", "account_id": "acc-1", "amount": {"minor": "lots", "currency": "EUR"}, "date": "2026-12-12", "description": "money well spent"}`, Transaction{Id: "123", AccountID: "acc-1", Amount: money.Money{Minor: 500, Currency: "EUR"}, Date: "2026-12-12", Description: "money well spent"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got Transaction
			err := json.Unmarshal([]byte(tc.json), &got)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected decode error, got none: %+v", got)
				}
				t.Logf("decode error: %v", err)
				return
			}
			if err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestTransaction_IsExpense(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		trans Transaction
		want  bool
	}{
		{"Below zero", Transaction{Amount: money.Money{Minor: -100, Currency: "EUR"}}, true},
		{"Above zero", Transaction{Amount: money.Money{Minor: 500, Currency: "EUR"}}, false},
		{"Zero", Transaction{Amount: money.Money{Minor: 0, Currency: "EUR"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.trans.IsExpense()

			if tt.want != got {
				t.Errorf("IsExpense() = %v, want %v", got, tt.want)
			}
		})
	}
}
