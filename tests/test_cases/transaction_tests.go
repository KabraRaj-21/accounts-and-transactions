//go:build slt

package test_cases

import (
	"net/http"
	"transaction/internal/entity"
	"transaction/tests/setup"
	"transaction/tests/util"

	"github.com/Swiggy/grill"
	"github.com/shopspring/decimal"
)

func GetTransactionTests(testEnv *setup.TestEnvironment) []grill.TestCase {
	tests := make([]grill.TestCase, 0)
	tests = append(tests, getTestsForTransactionRegistrationUseCase(testEnv)...)

	return tests
}

func getTestsForTransactionRegistrationUseCase(testEnv *setup.TestEnvironment) []grill.TestCase {
	validTransactionRequest := `{"account_id":"123","amount":10.5,"operation_type_id":1}`
	invalidTransactionRequest := `{"account_id":"123","amount":134.5}`
	withdrawlRequestExceedingBalance := `{"account_id":"123","amount":134.5,"operation_type_id":3}`

	return []grill.TestCase{
		{
			Name:  "Transaction_Registration_Request_Marshalling_Failure",
			Stubs: []grill.Stub{},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, RegisterTransactionUrl, "}") // invalid request

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusBadRequest),
			},
			Cleaners: []grill.Cleaner{},
		},
		{
			Name:  "Transaction_Registration_Request_Validation_Failure",
			Stubs: []grill.Stub{},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, RegisterTransactionUrl, invalidTransactionRequest)

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusBadRequest),
			},
			Cleaners: []grill.Cleaner{},
		},
		{
			Name:  "Transaction_Registration_DB_Failure",
			Stubs: []grill.Stub{},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, RegisterTransactionUrl, validTransactionRequest)

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusInternalServerError),
			},
			Cleaners: []grill.Cleaner{},
		},
		{
			Name: "Transaction_Registration_Account_Id_Not_Found",
			Stubs: []grill.Stub{
				testEnv.SetupTables(),
			},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, RegisterTransactionUrl, validTransactionRequest)

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusNotFound),
			},
			Cleaners: []grill.Cleaner{
				testEnv.CleanTables(),
			},
		},
		{
			Name: "Transaction_Registration_Balance_Valiation_Failure",
			Stubs: []grill.Stub{
				testEnv.SetupTables(),
				testEnv.MySQL.SeedFromCSVFile(AccountsTableName, "test_data/accounts.csv"),
			},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, RegisterTransactionUrl, withdrawlRequestExceedingBalance)

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusBadRequest),
			},
			Cleaners: []grill.Cleaner{
				testEnv.CleanTables(),
			},
		},
		{
			Name: "Transaction_Registration_Success",
			Stubs: []grill.Stub{
				testEnv.SetupTables(),
				testEnv.MySQL.SeedFromCSVFile(AccountsTableName, "test_data/accounts.csv"),
			},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, RegisterTransactionUrl, validTransactionRequest)
				var resBody entity.Transaction
				util.ExtractResponseBody(res, &resBody)

				return grill.ActionOutput(res.StatusCode, resBody.AccountId, resBody.OperationType, resBody.Amount.Equal(decimal.NewFromFloat(10.5)))
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusOK, "123", entity.OperationType_NORMAL_PURCHASE, true),
				testEnv.MySQL.AssertCount(TransactionsTableName, 1),
				testEnv.MySQL.AssertCount(AccountsTableName, 2),
			},
			Cleaners: []grill.Cleaner{
				testEnv.CleanTables(),
			},
		},
	}
}
