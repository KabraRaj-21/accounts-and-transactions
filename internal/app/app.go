package app

import (
	"accounts-and-transactions/internal/account"
	"accounts-and-transactions/internal/config"
	"accounts-and-transactions/internal/logger"
	"accounts-and-transactions/internal/repository"
	"accounts-and-transactions/internal/server/http"
	"accounts-and-transactions/internal/server/http/handler"
	"accounts-and-transactions/internal/transaction"
	validator_types "accounts-and-transactions/internal/transaction/types/validator_service"
	"accounts-and-transactions/internal/validator_service"
	"accounts-and-transactions/internal/validator_service/validator"
	"context"
	"sync"
)

type App struct {
	config *config.AppConfig
	server *http.Server
}

var (
	app     App
	onceApp sync.Once
)

func New(ctx context.Context, appConfig *config.AppConfig) *App {
	onceApp.Do(func() {
		// setup infra
		accountsDB, err := repository.NewAccountsSQLDB(ctx, appConfig.DBConfig)
		if err != nil {
			logger.WithContext(ctx).Errorf("error creating accounts db, err: %v", err)
			panic(err)
		}

		// setup repo
		repo := repository.NewMySQLRepository(accountsDB)

		// setup services
		accountService := account.NewAccountService(repo)
		validatorService := setupValidators()
		transactionService := transaction.NewTransactionService(repo, accountService, validatorService)

		// setup handlers
		accountsHander := handler.NewAccountHandler(accountService)
		transactionHandler := handler.NewTransactionHandler(transactionService)

		httpserver := http.NewServer(ctx, appConfig.HttpServerConfig, accountsHander, transactionHandler)

		app = App{
			config: appConfig,
			server: httpserver,
		}
	})
	return &app
}

func (app *App) Run(ctx context.Context) {
	app.server.Start(ctx)
}

func setupValidators() *validator_service.Service {
	balanceValidator := validator.NewAccountBalanceValidator()
	validatorService := validator_service.NewValidatorService()
	validatorService.RegisterValidator(validator_types.ActionType_TransactionRegistration, balanceValidator)

	return validatorService
}
