package account_service

import (
	"accounts-and-transactions/internal/entity"
	"context"
)

type AccountService interface {
	Create(ctx context.Context, req *CreateAccountRequest) (*entity.Account, error)
	Get(ctx context.Context, req *GetAccountRequest) (*entity.Account, error)
}
