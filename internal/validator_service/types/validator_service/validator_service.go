package validator_service

import (
	vs "accounts-and-transactions/internal/transaction/types/validator_service"
	"accounts-and-transactions/internal/validator_service/types/validator"
	"context"
)

type ValidatorService interface {
	RegisterValidator(actionType vs.ActionType, validator validator.Validator) error
	PerformValidation(ctx context.Context, req *vs.ValidationRequest) error
}
