//go:build slt

package test_cases

import (
	"fmt"
	"net/http"
	"transaction/internal/entity"
	"transaction/tests/setup"
	"transaction/tests/util"

	"github.com/Swiggy/grill"
	"github.com/shopspring/decimal"
)

func GetAccountTests(testEnv *setup.TestEnvironment) []grill.TestCase {
	tests := make([]grill.TestCase, 0)
	tests = append(tests, getTestsForAccountCreationUseCase(testEnv)...)
	tests = append(tests, getTestsForGetCreationUseCase(testEnv)...)

	return tests
}

func getTestsForAccountCreationUseCase(testEnv *setup.TestEnvironment) []grill.TestCase {
	validAccountCreationRequest := `{"document_number":"1234"}`
	return []grill.TestCase{
		{
			Name:  "Create_Account_Validation_Failure",
			Stubs: []grill.Stub{},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, CreateAccountUrl, "}") // invalid request

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusBadRequest),
			},
			Cleaners: []grill.Cleaner{},
		},
		{
			Name:  "Create_Account_DB_Failure",
			Stubs: []grill.Stub{},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, CreateAccountUrl, validAccountCreationRequest)

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusInternalServerError),
			},
			Cleaners: []grill.Cleaner{},
		},
		{
			Name: "Create_Account_Success",
			Stubs: []grill.Stub{
				testEnv.SetupTables(),
			},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodPost, CreateAccountUrl, validAccountCreationRequest)
				var resBody entity.Account
				util.ExtractResponseBody(res, &resBody)

				return grill.ActionOutput(res.StatusCode, resBody.DocumentNumber, resBody.Balance.Equal(decimal.Zero))
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusOK, "1234", true),
				testEnv.MySQL.AssertCount(AccountsTableName, 1),
			},
			Cleaners: []grill.Cleaner{
				testEnv.CleanTables(),
			},
		},
	}
}

func getTestsForGetCreationUseCase(testEnv *setup.TestEnvironment) []grill.TestCase {
	testAccountId := "123"
	return []grill.TestCase{
		{
			Name:  "Get_Account_DB_Failure",
			Stubs: []grill.Stub{},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodGet, fmt.Sprintf(GetAccountUrl, testAccountId), "")

				return grill.ActionOutput(res.StatusCode)
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusInternalServerError),
			},
			Cleaners: []grill.Cleaner{},
		},
		{
			Name: "Get_Account_Not_Found",
			Stubs: []grill.Stub{
				testEnv.SetupTables(),
			},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodGet, fmt.Sprintf(GetAccountUrl, testAccountId), "")

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
			Name: "Get_Account_Success",
			Stubs: []grill.Stub{
				testEnv.SetupTables(),
				testEnv.MySQL.SeedFromCSVFile(AccountsTableName, "test_data/accounts.csv"),
			},
			Action: func() interface{} {
				res := util.MakeHTTPApiCall(http.MethodGet, fmt.Sprintf(GetAccountUrl, testAccountId), "")
				var resBody entity.Account
				util.ExtractResponseBody(res, &resBody)

				return grill.ActionOutput(res.StatusCode, resBody.DocumentNumber, resBody.Balance.Equal(decimal.NewFromFloat(100.56)))
			},
			Assertions: []grill.Assertion{
				grill.AssertOutput(http.StatusOK, "777", true),
			},
			Cleaners: []grill.Cleaner{
				testEnv.CleanTables(),
			},
		},
	}
}
