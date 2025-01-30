package account

import (
	"context"
	"transaction/internal/account/types/account_repository"
	"transaction/internal/account/types/account_service"
	"transaction/internal/entity"

	"github.com/shopspring/decimal"
)

type Service struct {
	repository account_repository.AccountRepository
}

func NewAccountService(repository account_repository.AccountRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (as *Service) Create(ctx context.Context, req *account_service.CreateAccountRequest) (*entity.Account, error) {
	newAccount := &entity.Account{
		DocumentNumber: req.DocumentNumber,
		Balance:        decimal.Zero,
	}
	return as.repository.CreateAccount(ctx, newAccount)
}

func (as *Service) Get(ctx context.Context, req *account_service.GetAccountRequest) (*entity.Account, error) {
	return as.repository.GetAccountById(ctx, req.Id)
}
