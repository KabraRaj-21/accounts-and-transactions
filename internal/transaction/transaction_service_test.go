//go:build unit

package transaction

import (
	"accounts-and-transactions/internal/account/types/account_service"
	"accounts-and-transactions/internal/entity"
	"accounts-and-transactions/internal/transaction/mocks"
	"accounts-and-transactions/internal/transaction/types/transaction_service"
	"accounts-and-transactions/internal/transaction/types/validator_service"
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionService_RegisterTransaction(t *testing.T) {
	testCtx := context.Background()
	testRequest := &transaction_service.RegisterTransactionRequest{
		AccountId:     "1234",
		Amount:        456.78,
		OperationType: 1,
	}
	testAccount := &entity.Account{
		Id:             "1234",
		DocumentNumber: "5678",
		Balance:        decimal.NewFromFloat(1000.5),
	}

	accountInfoRequestValidator := func(r *account_service.GetAccountRequest) bool {
		return testRequest.AccountId == r.Id
	}
	validatorServiceRequestValidator := func(r *validator_service.ValidationRequest) bool {
		return testRequest.AccountId == r.Transaction.AccountId &&
			r.Transaction.Amount.Equal(decimal.NewFromFloat(testRequest.Amount)) &&
			testRequest.OperationType == int(r.Transaction.OperationType) &&
			testAccount == r.Account
	}
	repositoryRequestValidator := func(r *entity.Transaction) bool {
		return testRequest.AccountId == r.AccountId &&
			r.Amount.Equal(decimal.NewFromFloat(testRequest.Amount)) &&
			testRequest.OperationType == int(r.OperationType)
	}

	testCases := map[string]struct {
		request                     *transaction_service.RegisterTransactionRequest
		accountInfoServiceBehaviour func(m *mocks.AccountInfoService)
		validatorServiceBehaviour   func(m *mocks.ValidatorService)
		repositoryBehaviour         func(m *mocks.TransactionRepository)
		expectedResponse            *entity.Transaction
		isErrorExpected             bool
	}{
		"invalid request, should return error": {
			request:         &transaction_service.RegisterTransactionRequest{OperationType: 5},
			isErrorExpected: true,
		},

		"failed to get account details, should return error": {
			request: testRequest,
			accountInfoServiceBehaviour: func(m *mocks.AccountInfoService) {
				m.On("Get", testCtx, mock.MatchedBy(accountInfoRequestValidator)).Return(nil, errors.New("error getting account"))
			},
			isErrorExpected: true,
		},

		"transaction validation failure, should return error": {
			request: testRequest,

			accountInfoServiceBehaviour: func(m *mocks.AccountInfoService) {
				m.On("Get", testCtx, mock.MatchedBy(accountInfoRequestValidator)).Return(testAccount, nil)
			},

			validatorServiceBehaviour: func(m *mocks.ValidatorService) {
				m.On("PerformValidation", testCtx, mock.MatchedBy(validatorServiceRequestValidator)).Return(errors.New("validation failures"))
			},
			isErrorExpected: true,
		},

		"repository save operation failure, should return error": {
			request: testRequest,

			accountInfoServiceBehaviour: func(m *mocks.AccountInfoService) {
				m.On("Get", testCtx, mock.MatchedBy(accountInfoRequestValidator)).Return(testAccount, nil)
			},

			validatorServiceBehaviour: func(m *mocks.ValidatorService) {
				m.On("PerformValidation", testCtx, mock.MatchedBy(validatorServiceRequestValidator)).Return(nil)
			},

			repositoryBehaviour: func(m *mocks.TransactionRepository) {
				m.On("RegisterTransaction", testCtx, testAccount, mock.MatchedBy(repositoryRequestValidator)).Return(nil, errors.New("error saving transaction"))
			},
			isErrorExpected: true,
		},

		"no failures, should return proper entity": {
			request: testRequest,

			accountInfoServiceBehaviour: func(m *mocks.AccountInfoService) {
				m.On("Get", testCtx, mock.MatchedBy(accountInfoRequestValidator)).Return(testAccount, nil)
			},

			validatorServiceBehaviour: func(m *mocks.ValidatorService) {
				m.On("PerformValidation", testCtx, mock.MatchedBy(validatorServiceRequestValidator)).Return(nil)
			},

			repositoryBehaviour: func(m *mocks.TransactionRepository) {
				m.On("RegisterTransaction", testCtx, testAccount, mock.MatchedBy(repositoryRequestValidator)).
					Return(&entity.Transaction{Id: "888", AccountId: testRequest.AccountId,
						OperationType: entity.OperationType(testRequest.OperationType), Amount: decimal.NewFromFloat(testRequest.Amount)}, nil)
			},
			isErrorExpected: false,
			expectedResponse: &entity.Transaction{Id: "888", AccountId: testRequest.AccountId,
				OperationType: entity.OperationType(testRequest.OperationType), Amount: decimal.NewFromFloat(testRequest.Amount)},
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			accountInfoMock := mocks.NewAccountInfoService(t)
			validatorMock := mocks.NewValidatorService(t)
			repositoryMock := mocks.NewTransactionRepository(t)

			if tc.accountInfoServiceBehaviour != nil {
				tc.accountInfoServiceBehaviour(accountInfoMock)
			}

			if tc.validatorServiceBehaviour != nil {
				tc.validatorServiceBehaviour(validatorMock)
			}

			if tc.repositoryBehaviour != nil {
				tc.repositoryBehaviour(repositoryMock)
			}

			transactionService := NewTransactionService(repositoryMock, accountInfoMock, validatorMock)
			res, err := transactionService.RegisterTransaction(testCtx, tc.request)

			if tc.isErrorExpected {
				assert.True(t, err != nil)
			} else {
				assert.True(t, err == nil)
				assert.Equal(t, tc.request.AccountId, res.AccountId)
				assert.Equal(t, decimal.NewFromFloat(tc.request.Amount), res.Amount)
				assert.Equal(t, tc.request.OperationType, int(res.OperationType))
				assert.True(t, res.Id != "")
			}
			accountInfoMock.AssertExpectations(t)
			validatorMock.AssertExpectations(t)
			repositoryMock.AssertExpectations(t)
		})
	}
}
