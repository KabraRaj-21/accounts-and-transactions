package validator

import (
	"accounts-and-transactions/internal/transaction/types/validator_service"
	"context"
)

type Validator interface {
	Validate(ctx context.Context, req *validator_service.ValidationRequest) error
}
