//go:build unit

package handler

import (
	"accounts-and-transactions/internal/entity"
	"accounts-and-transactions/internal/errors/tserror"
	"accounts-and-transactions/internal/transaction/mocks"
	"accounts-and-transactions/internal/transaction/types/transaction_service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionHandler_RegisterTransaction(t *testing.T) {
	testCases := map[string]struct {
		request                  *transaction_service.RegisterTransactionRequest
		isServiceCalled          bool
		serviceResponse          *entity.Transaction
		serviceError             error
		isResponseExpected       bool
		expectedHttpResponseCode int
		expectedResponse         *entity.Transaction
	}{
		"validation failure, should return error": {
			request:                  nil,
			isServiceCalled:          false,
			isResponseExpected:       false,
			expectedHttpResponseCode: http.StatusBadRequest,
		},
		"service failure, should return error": {
			request:                  &transaction_service.RegisterTransactionRequest{AccountId: "123", Amount: 100.5, OperationType: 1},
			isServiceCalled:          true,
			serviceError:             tserror.New(tserror.ErrorType_OPERATION_FAILURE, "registration failure"),
			isResponseExpected:       false,
			expectedHttpResponseCode: http.StatusInternalServerError,
		},
		"success": {
			request:                  &transaction_service.RegisterTransactionRequest{AccountId: "123", Amount: 100.5, OperationType: 1},
			isServiceCalled:          true,
			serviceResponse:          &entity.Transaction{Id: "007", AccountId: "123", Amount: decimal.NewFromFloat(100.5), OperationType: entity.OperationType_NORMAL_PURCHASE},
			serviceError:             nil,
			isResponseExpected:       true,
			expectedHttpResponseCode: http.StatusOK,
			expectedResponse:         &entity.Transaction{Id: "007", AccountId: "123", Amount: decimal.NewFromFloat(100.5), OperationType: entity.OperationType_NORMAL_PURCHASE},
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			transactionService := mocks.NewTransactionService(t)
			if tc.isServiceCalled {
				transactionService.On("RegisterTransaction", mock.Anything, mock.Anything).Return(tc.serviceResponse, tc.serviceError)
			}

			transactionHandler := NewTransactionHandler(transactionService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = getHttpRequest(http.MethodPost, "api/v1/transactions", tc.request)

			transactionHandler.RegisterTransaction(c)

			assert.Equal(t, tc.expectedHttpResponseCode, w.Code)

			if tc.isResponseExpected {
				var accountResponse *entity.Transaction
				parseHttpRespose(w, &accountResponse)

				assert.Equal(t, tc.expectedResponse.Id, accountResponse.Id)
				assert.Equal(t, tc.expectedResponse.AccountId, accountResponse.AccountId)
				assert.True(t, tc.expectedResponse.Amount.Equal(accountResponse.Amount))
				assert.Equal(t, tc.expectedResponse.OperationType, accountResponse.OperationType)
			}
		})
	}
}
