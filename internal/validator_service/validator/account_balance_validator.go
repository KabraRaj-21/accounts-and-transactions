package validator

import (
	"accounts-and-transactions/internal/errors/tserror"
	"accounts-and-transactions/internal/transaction/types/validator_service"
	"context"

	"github.com/shopspring/decimal"
)

type AccountBalanceValidator struct {
	// todo: take minimum balance as a config here
}

func NewAccountBalanceValidator() *AccountBalanceValidator {
	return &AccountBalanceValidator{}
}

func (v *AccountBalanceValidator) Validate(ctx context.Context, req *validator_service.ValidationRequest) error {
	if req.Transaction.IsCreditTransaction() {
		return nil // no limit of amount to be deposited
	}

	hasEnoughBalance := req.Account.Balance.GreaterThanOrEqual(req.Transaction.Amount.Add(decimal.NewFromFloat(0))) // assuming minimum balance to be zero here
	if !hasEnoughBalance {
		return tserror.New(tserror.ErrorType_INVALID_REQUEST, "account does not have enough balance for this transaction")
	}
	return nil
}
