package account_repository

import (
	"context"
	"transaction/internal/entity"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, req *entity.Account) (*entity.Account, error)
	GetAccountById(ctx context.Context, id string) (*entity.Account, error)
}
