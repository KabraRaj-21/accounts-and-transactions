//go:build unit

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction/internal/account/mocks"
	"transaction/internal/account/types/account_service"
	"transaction/internal/entity"
	"transaction/internal/errors/tserror"
	"transaction/internal/server/http/types"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountHandler_CreateAccount(t *testing.T) {
	testCases := map[string]struct {
		request                  *account_service.CreateAccountRequest
		isServiceCalled          bool
		serviceResponse          *entity.Account
		serviceError             error
		isResponseExpected       bool
		expectedHttpResponseCode int
		expectedResponse         *entity.Account
	}{
		"validation failure, should return error": {
			request:                  nil,
			isServiceCalled:          false,
			isResponseExpected:       false,
			expectedHttpResponseCode: http.StatusBadRequest,
		},
		"service failure, should return error": {
			request:                  &account_service.CreateAccountRequest{DocumentNumber: "123"},
			isServiceCalled:          true,
			serviceError:             tserror.New(tserror.ErrorType_OPERATION_FAILURE, "creation failure"),
			isResponseExpected:       false,
			expectedHttpResponseCode: http.StatusInternalServerError,
		},
		"success": {
			request:                  &account_service.CreateAccountRequest{DocumentNumber: "123"},
			isServiceCalled:          true,
			serviceResponse:          &entity.Account{Id: "007", DocumentNumber: "123"},
			serviceError:             nil,
			isResponseExpected:       true,
			expectedHttpResponseCode: http.StatusOK,
			expectedResponse:         &entity.Account{Id: "007", DocumentNumber: "123"},
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			accountService := mocks.NewAccountService(t)
			if tc.isServiceCalled {
				accountService.On("Create", mock.Anything, mock.Anything).Return(tc.serviceResponse, tc.serviceError)
			}

			accountHandler := NewAccountHandler(accountService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = getHttpRequest(http.MethodPost, "api/v1/accounts", tc.request)

			accountHandler.CreateAccount(c)

			assert.Equal(t, tc.expectedHttpResponseCode, w.Code)

			if tc.isResponseExpected {
				var accountResponse *entity.Account
				parseHttpRespose(w, &accountResponse)

				assert.Equal(t, tc.expectedResponse.Id, accountResponse.Id)
				assert.Equal(t, tc.expectedResponse.DocumentNumber, accountResponse.DocumentNumber)
				assert.True(t, tc.expectedResponse.Balance.Equal(accountResponse.Balance))
			}
		})
	}
}

func TestAccountHandler_GetAccount(t *testing.T) {
	testCases := map[string]struct {
		requestParam             string
		isServiceCalled          bool
		serviceResponse          *entity.Account
		serviceError             error
		isResponseExpected       bool
		expectedHttpResponseCode int
		expectedResponse         *entity.Account
	}{
		"validation failure, should return error": {
			requestParam:             "",
			isServiceCalled:          false,
			isResponseExpected:       false,
			expectedHttpResponseCode: http.StatusBadRequest,
		},
		"service failure, should return error": {
			requestParam:             "123",
			isServiceCalled:          true,
			serviceError:             tserror.New(tserror.ErrorType_NOT_FOUND, "not found"),
			isResponseExpected:       false,
			expectedHttpResponseCode: http.StatusNotFound,
		},
		"success": {
			requestParam:             "007",
			isServiceCalled:          true,
			serviceResponse:          &entity.Account{Id: "007", DocumentNumber: "123"},
			serviceError:             nil,
			isResponseExpected:       true,
			expectedHttpResponseCode: http.StatusOK,
			expectedResponse:         &entity.Account{Id: "007", DocumentNumber: "123"},
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			accountService := mocks.NewAccountService(t)
			if tc.isServiceCalled {
				accountService.On("Get", mock.Anything, mock.Anything).Return(tc.serviceResponse, tc.serviceError)
			}

			accountHandler := NewAccountHandler(accountService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = getHttpRequest(http.MethodGet, "api/v1/accounts", nil)
			c.Params = []gin.Param{
				{
					Key:   "accountId",
					Value: tc.requestParam,
				},
			}

			accountHandler.GetAccount(c)

			assert.Equal(t, tc.expectedHttpResponseCode, w.Code)

			if tc.isResponseExpected {
				var accountResponse *entity.Account
				parseHttpRespose(w, &accountResponse)

				assert.Equal(t, tc.expectedResponse.Id, accountResponse.Id)
				assert.Equal(t, tc.expectedResponse.DocumentNumber, accountResponse.DocumentNumber)
				assert.True(t, tc.expectedResponse.Balance.Equal(accountResponse.Balance))
			}
		})
	}
}

func getHttpRequest(method, endpoint string, requestBody interface{}) *http.Request {
	if requestBody != nil {
		marshalledBody, _ := json.Marshal(requestBody)
		body := bytes.NewBuffer(marshalledBody)
		req, _ := http.NewRequest(method, endpoint, body)
		return req
	}
	req, _ := http.NewRequest(method, endpoint, nil)
	return req

}

func parseHttpRespose(w *httptest.ResponseRecorder, res interface{}) {
	var apiResponse types.APIResponse
	json.Unmarshal(w.Body.Bytes(), &apiResponse)

	dataBytes, _ := json.Marshal(apiResponse.Data)
	json.Unmarshal(dataBytes, &res)
}
