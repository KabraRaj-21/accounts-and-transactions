package account_repository

import (
	"accounts-and-transactions/internal/entity"
	"context"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, req *entity.Account) (*entity.Account, error)
	GetAccountById(ctx context.Context, id string) (*entity.Account, error)
}
