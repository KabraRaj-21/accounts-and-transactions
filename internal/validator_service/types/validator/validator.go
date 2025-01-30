package validator

import (
	"context"
	"transaction/internal/transaction/types/validator_service"
)

type Validator interface {
	Validate(ctx context.Context, req *validator_service.ValidationRequest) error
}
