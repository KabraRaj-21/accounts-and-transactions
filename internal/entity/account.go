package entity

import "github.com/shopspring/decimal"

type Account struct {
	Id             string
	DocumentNumber string
	Balance        decimal.Decimal
}

func (a *Account) UpdateBalance(change decimal.Decimal) {
	a.Balance = a.Balance.Add(change)
}
