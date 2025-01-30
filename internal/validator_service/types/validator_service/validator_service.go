package validator_service

import (
	"context"
	vs "transaction/internal/transaction/types/validator_service"
	"transaction/internal/validator_service/types/validator"
)

type ValidatorService interface {
	RegisterValidator(actionType vs.ActionType, validator validator.Validator) error
	PerformValidation(ctx context.Context, req *vs.ValidationRequest) error
}
