package validator_service

import (
	"context"
	vs "transaction/internal/transaction/types/validator_service"
	"transaction/internal/validator_service/types/validator"
)

type Service struct {
	validatorMap map[vs.ActionType][]validator.Validator
}

func NewValidatorService() *Service {
	return &Service{
		validatorMap: make(map[vs.ActionType][]validator.Validator, 0),
	}
}

func (s *Service) RegisterValidator(actionType vs.ActionType, v validator.Validator) error {
	s.validatorMap[actionType] = append(s.validatorMap[actionType], v)
	return nil
}

func (s *Service) PerformValidation(ctx context.Context, req *vs.ValidationRequest) error {
	actionType := req.ActionType
	for _, v := range s.validatorMap[actionType] {
		err := v.Validate(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
