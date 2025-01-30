//go:build slt

package tests

import (
	"context"
	"testing"
	"transaction/internal/app"
	"transaction/internal/config"
	"transaction/tests/setup"
	"transaction/tests/test_cases"

	"github.com/Swiggy/grill"
)

func TestFunctional(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional tests")
	}

	// setup env
	testEnv := setup.NewTestEnvironment()
	if err := testEnv.StartAll(); err != nil {
		t.Errorf("error creating test env: %v", err)
	}
	defer testEnv.StopAll()
	setup.SetupEnvironmentVariables(testEnv)

	// start application
	testCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appConfig := config.GetAppConfig(testCtx)
	application := app.New(testCtx, appConfig)
	go application.Run(testCtx)

	t.Run("Account SLTs", func(t *testing.T) {
		grill.Run(t, test_cases.GetAccountTests(testEnv))
	})

	t.Run("Transaction SLTs", func(t *testing.T) {
		grill.Run(t, test_cases.GetTransactionTests(testEnv))
	})

}
