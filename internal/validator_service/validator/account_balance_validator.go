package validator

import (
	"context"
	"transaction/internal/errors/tserror"
	"transaction/internal/transaction/types/validator_service"

	"github.com/shopspring/decimal"
)

type AccountBalanceValidator struct {
	// todo: take minimum balance as a config here
}

func NewAccountBalanceValidator() *AccountBalanceValidator {
	return &AccountBalanceValidator{}
}

func (v *AccountBalanceValidator) Validate(ctx context.Context, req *validator_service.ValidationRequest) error {
	hasEnoughBalance := req.Account.Balance.GreaterThanOrEqual(req.Transaction.Amount.Add(decimal.NewFromFloat(0))) // assuming minimum balance to be zero here
	if !hasEnoughBalance {
		return tserror.New(tserror.ErrorType_INVALID_REQUEST, "account does not have enough balance for this transaction")
	}
	return nil
}
