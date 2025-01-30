//go:build unit

package validator_service

import (
	"accounts-and-transactions/internal/transaction/types/validator_service"
	"accounts-and-transactions/internal/validator_service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationService_RegisterValidator(t *testing.T) {
	validatorService := NewValidatorService()
	err := validatorService.RegisterValidator(validator_service.ActionType_TransactionRegistration, mocks.NewValidator(t))
	assert.Nil(t, err)
}

func TestValidationService_PerformValidation(t *testing.T) {
	testActionType := validator_service.ActionType_TransactionRegistration
	testCtx := context.Background()

	testCases := map[string]struct {
		input                   *validator_service.ValidationRequest
		isValidatorCallExpected bool
		validatorError          error
		isErrorExpected         bool
	}{
		"no error": {
			input:                   &validator_service.ValidationRequest{ActionType: validator_service.ActionType_TransactionRegistration},
			isValidatorCallExpected: true,
			validatorError:          nil,
			isErrorExpected:         false,
		},
		"validation failure": {
			input:                   &validator_service.ValidationRequest{ActionType: validator_service.ActionType_TransactionRegistration},
			isValidatorCallExpected: true,
			validatorError:          errors.New("validation failure"),
			isErrorExpected:         true,
		},
		"no validator found": {
			input:                   &validator_service.ValidationRequest{ActionType: validator_service.ActionType_Unknown},
			isValidatorCallExpected: false,
			validatorError:          nil,
			isErrorExpected:         false,
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			validator := mocks.NewValidator(t)

			if tc.isValidatorCallExpected {
				validator.On("Validate", testCtx, tc.input).Return(tc.validatorError)
			}

			validatorService := NewValidatorService()
			validatorService.RegisterValidator(testActionType, validator)

			err := validatorService.PerformValidation(testCtx, tc.input)
			assert.Equal(t, tc.isErrorExpected, err != nil)
			validator.AssertExpectations(t)
		})
	}
}
