//go:build unit

package entity

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsCreditTransaction(t *testing.T) {
	testTransaction := &Transaction{
		OperationType: OperationType_CREDIT_VOUCHER,
	}

	assert.True(t, testTransaction.IsCreditTransaction())

	testTransaction = &Transaction{
		OperationType: OperationType_PURCHASE_WITH_INSTALLMENTS,
	}
	assert.False(t, testTransaction.IsCreditTransaction())
}

func TestTransaction_GetBalanceChange(t *testing.T) {
	testTransaction := &Transaction{
		OperationType: OperationType_CREDIT_VOUCHER,
		Amount:        decimal.NewFromFloat(100.50),
	}

	assert.True(t, testTransaction.GetBalanceChange().Equal(decimal.NewFromFloat(100.50)))

	testTransaction = &Transaction{
		OperationType: OperationType_WITHDRAWL,
		Amount:        decimal.NewFromFloat(25.50),
	}

	decimalVal := testTransaction.GetBalanceChange()
	val, exact := decimalVal.Float64()
	assert.True(t, exact)
	assert.Equal(t, -25.50, val)
}
