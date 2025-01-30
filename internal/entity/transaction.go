package entity

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type OperationType int

const (
	OperationType_NORMAL_PURCHASE            OperationType = 1
	OperationType_PURCHASE_WITH_INSTALLMENTS OperationType = 2
	OperationType_WITHDRAWL                  OperationType = 3
	OperationType_CREDIT_VOUCHER             OperationType = 4
)

func ParseOperationTypeFromInt(k int) (OperationType, error) {
	if k < int(OperationType_NORMAL_PURCHASE) || k > int(OperationType_CREDIT_VOUCHER) {
		return 0, errors.New("invalid ordertion type")
	}
	return OperationType(k), nil
}

type Transaction struct {
	Id            string
	OperationType OperationType
	AccountId     string
	Amount        decimal.Decimal
	Timestamp     time.Time
}

func (t *Transaction) GetBalanceChange() decimal.Decimal {
	if t.IsCreditTransaction() {
		return t.Amount
	}
	return t.Amount.Neg()
}

func (t *Transaction) IsCreditTransaction() bool {
	return t.OperationType == OperationType_CREDIT_VOUCHER
}
