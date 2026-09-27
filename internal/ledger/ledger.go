package ledger

import (
	"github.com/mvoss02/munto/internal/money"
)

type Transaction struct {
	Id          string      `json:"id"`
	AccountID   string      `json:"account_id"`
	Amount      money.Money `json:"amount"`
	Date        string      `json:"date"` // Type will be repalced
	Description string      `json:"description"`
}

type Account struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Currecny string `json:"currency"`
}

func (t Transaction) IsExpense() bool {
	return t.Amount.Minor < 0
}
