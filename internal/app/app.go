package app

import (
	"context"
	"sync"
	"transaction/internal/account"
	"transaction/internal/config"
	"transaction/internal/logger"
	"transaction/internal/repository"
	"transaction/internal/server/http"
	"transaction/internal/server/http/handler"
	"transaction/internal/transaction"
	validator_types "transaction/internal/transaction/types/validator_service"
	"transaction/internal/validator_service"
	"transaction/internal/validator_service/validator"
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
