//go:build unit

package entity

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAccount_UpdateBalance(t *testing.T) {
	testCases := map[string]struct {
		change          decimal.Decimal
		initialBalance  decimal.Decimal
		expectedBalance decimal.Decimal
	}{
		"Positive change": {
			change:          decimal.NewFromFloat(100.56),
			initialBalance:  decimal.NewFromFloat(99.44),
			expectedBalance: decimal.NewFromFloat(200),
		},
		"Negative change": {
			change:          decimal.NewFromFloat(100.56),
			initialBalance:  decimal.NewFromFloat(-99.44),
			expectedBalance: decimal.NewFromFloat(1.12),
		},
		"No change": {
			change:          decimal.NewFromFloat(100.56),
			initialBalance:  decimal.NewFromFloat(0),
			expectedBalance: decimal.NewFromFloat(100.56),
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			testAccount := &Account{
				Id:             "123",
				DocumentNumber: "456",
				Balance:        tc.initialBalance,
			}

			testAccount.UpdateBalance(tc.change)

			assert.True(t, testAccount.Balance.Equal(tc.expectedBalance))
			assert.Equal(t, testAccount.Id, "123")
			assert.Equal(t, testAccount.DocumentNumber, "456")
		})
	}
}
