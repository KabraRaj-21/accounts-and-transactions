package validator_service

import (
	"context"
)

type ValidatorService interface {
	PerformValidation(ctx context.Context, req *ValidationRequest) error
}
