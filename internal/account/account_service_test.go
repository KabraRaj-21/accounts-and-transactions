//go:build unit

package account

import (
	"context"
	"errors"
	"testing"
	"transaction/internal/account/mocks"
	"transaction/internal/account/types/account_service"
	"transaction/internal/entity"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_Create(t *testing.T) {
	testReq := &account_service.CreateAccountRequest{
		DocumentNumber: "1234",
	}
	testCtx := context.Background()

	testCases := map[string]struct {
		request             *account_service.CreateAccountRequest
		repositoryBehaviour func(m *mocks.AccountRepository)
		expectedResponse    *entity.Account
		isErrorExpected     bool
	}{
		"repository operation failure, should return error": {
			request: testReq,
			repositoryBehaviour: func(m *mocks.AccountRepository) {
				m.On("CreateAccount", testCtx, mock.MatchedBy(func(r *entity.Account) bool {
					assert.Equal(t, testReq.DocumentNumber, r.DocumentNumber)
					assert.Equal(t, decimal.Zero, r.Balance)

					return true
				})).Return(nil, errors.New("creation failure"))
			},

			isErrorExpected: true,
		},

		"repository operation successful, should return account entity": {
			request: testReq,
			repositoryBehaviour: func(m *mocks.AccountRepository) {
				m.On("CreateAccount", testCtx, mock.MatchedBy(func(r *entity.Account) bool {
					assert.Equal(t, testReq.DocumentNumber, r.DocumentNumber)
					assert.Equal(t, decimal.Zero, r.Balance)

					return true
				})).Return(&entity.Account{Id: "testId", DocumentNumber: testReq.DocumentNumber, Balance: decimal.Zero}, nil)
			},

			isErrorExpected:  false,
			expectedResponse: &entity.Account{DocumentNumber: "1234", Balance: decimal.Zero},
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			mockRepository := mocks.NewAccountRepository(t)

			tc.repositoryBehaviour(mockRepository)
			accountService := NewAccountService(mockRepository)

			res, err := accountService.Create(testCtx, tc.request)
			if tc.isErrorExpected {
				assert.True(t, err != nil)
			} else {
				assert.True(t, err == nil)
				assert.Equal(t, tc.expectedResponse.DocumentNumber, res.DocumentNumber)
				assert.True(t, tc.expectedResponse.Balance.Equal(decimal.NewFromFloat(0.0)))
				assert.True(t, res.Id != "")
			}
		})
	}
}

func TestAccountService_Get(t *testing.T) {
	testReq := &account_service.GetAccountRequest{
		Id: "1234",
	}
	testCtx := context.Background()

	sampleAccountEntity := &entity.Account{
		Id:             "1234",
		DocumentNumber: "5678",
		Balance:        decimal.NewFromFloat(1234.56),
	}

	testCases := map[string]struct {
		request             *account_service.GetAccountRequest
		repositoryBehaviour func(m *mocks.AccountRepository)
		expectedResponse    *entity.Account
		isErrorExpected     bool
	}{
		"repository operation failure, should return error": {
			request: testReq,
			repositoryBehaviour: func(m *mocks.AccountRepository) {
				m.On("GetAccountById", testCtx, testReq.Id).Return(nil, errors.New("failed to get account"))
			},
			isErrorExpected: true,
		},

		"repository operation is successful, should return entity": {
			request: testReq,
			repositoryBehaviour: func(m *mocks.AccountRepository) {
				m.On("GetAccountById", testCtx, testReq.Id).Return(sampleAccountEntity, nil)
			},
			isErrorExpected:  false,
			expectedResponse: sampleAccountEntity,
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			mockRepository := mocks.NewAccountRepository(t)

			tc.repositoryBehaviour(mockRepository)
			accountService := NewAccountService(mockRepository)

			res, err := accountService.Get(testCtx, tc.request)
			if tc.isErrorExpected {
				assert.True(t, err != nil)
			} else {
				assert.True(t, err == nil)
				assert.Equal(t, tc.expectedResponse, res)
			}
		})
	}
}
