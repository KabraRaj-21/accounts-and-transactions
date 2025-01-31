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
		operationType     entity.OperationType
		accountBalance    decimal.Decimal
		transactionAmount decimal.Decimal
		isErrorExpected   bool
	}{
		"Non Debit transaction": {
			operationType:     entity.OperationType_CREDIT_VOUCHER,
			accountBalance:    decimal.NewFromFloat(1000),
			transactionAmount: decimal.NewFromFloat(1000.04),
			isErrorExpected:   false,
		},
		"Account does not have enough balance": {
			operationType:     entity.OperationType_NORMAL_PURCHASE,
			accountBalance:    decimal.NewFromFloat(1000),
			transactionAmount: decimal.NewFromFloat(1000.04),
			isErrorExpected:   true,
		},
		"Account has enough balance": {
			operationType:     entity.OperationType_PURCHASE_WITH_INSTALLMENTS,
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
					Amount:        tc.transactionAmount,
					OperationType: tc.operationType,
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
