package account_service

import (
	"context"
	"transaction/internal/entity"
)

type AccountService interface {
	Create(ctx context.Context, req *CreateAccountRequest) (*entity.Account, error)
	Get(ctx context.Context, req *GetAccountRequest) (*entity.Account, error)
}
