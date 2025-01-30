package transaction

import (
	"context"
	"time"
	"transaction/internal/account/types/account_service"
	"transaction/internal/entity"
	"transaction/internal/errors/tserror"
	account_info "transaction/internal/transaction/types/account_info_service"
	repo "transaction/internal/transaction/types/transaction_repository"
	"transaction/internal/transaction/types/transaction_service"
	"transaction/internal/transaction/types/validator_service"

	"github.com/shopspring/decimal"
)

type Service struct {
	repository         repo.TransactionRepository
	validatorService   validator_service.ValidatorService
	accountInfoService account_info.AccountInfoService
}

func NewTransactionService(repository repo.TransactionRepository,
	accountInfoService account_info.AccountInfoService,
	validatorService validator_service.ValidatorService) *Service {
	return &Service{
		repository:         repository,
		accountInfoService: accountInfoService,
		validatorService:   validatorService,
	}
}

func (ts *Service) RegisterTransaction(ctx context.Context, req *transaction_service.RegisterTransactionRequest) (*entity.Transaction, error) {
	operationType, err := entity.ParseOperationTypeFromInt(req.OperationType)
	if err != nil {
		return nil, tserror.Wrap(tserror.ErrorType_INVALID_REQUEST, "opeartion type is not valid", err)
	}
	transaction := &entity.Transaction{
		OperationType: operationType,
		Amount:        decimal.NewFromFloat(req.Amount),
		AccountId:     req.AccountId,
		Timestamp:     time.Now(),
	}

	account, err := ts.accountInfoService.Get(ctx, &account_service.GetAccountRequest{Id: req.AccountId})
	if err != nil {
		return nil, err
	}

	validationRequest := &validator_service.ValidationRequest{
		ActionType:  validator_service.ActionType_TransactionRegistration,
		Account:     account,
		Transaction: transaction,
	}
	err = ts.validatorService.PerformValidation(ctx, validationRequest)
	if err != nil {
		return nil, err
	}

	account.UpdateBalance(transaction.GetBalanceChange())

	return ts.repository.RegisterTransaction(ctx, account, transaction)
}
