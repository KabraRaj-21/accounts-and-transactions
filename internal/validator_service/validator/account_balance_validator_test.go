//go:build unit

package validator

import (
	"accounts-and-transactions/internal/entity"
	"accounts-and-transactions/internal/transaction/types/validator_service"
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAccountBalanceValidator_Validate(t *testing.T) {
	testCases := map[string]struct {
		accountBalance    decimal.Decimal
		transactionAmount decimal.Decimal
		isErrorExpected   bool
	}{
		"Account does not have enough balance": {
			accountBalance:    decimal.NewFromFloat(1000),
			transactionAmount: decimal.NewFromFloat(1000.04),
			isErrorExpected:   true,
		},
		"Account has enough balance": {
			accountBalance:    decimal.NewFromFloat(1000),
			transactionAmount: decimal.NewFromFloat(500.04),
			isErrorExpected:   false,
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			accountBalanceValidator := NewAccountBalanceValidator()
			req := &validator_service.ValidationRequest{
				Transaction: &entity.Transaction{
					Amount: tc.transactionAmount,
				},
				Account: &entity.Account{
					Balance: tc.accountBalance,
				},
			}

			err := accountBalanceValidator.Validate(context.Background(), req)
			assert.Equal(t, tc.isErrorExpected, err != nil)
		})
	}
}
