package account_info_service

import (
	"accounts-and-transactions/internal/account/types/account_service"
	"accounts-and-transactions/internal/entity"
	"context"
)

type AccountInfoService interface {
	Get(ctx context.Context, req *account_service.GetAccountRequest) (*entity.Account, error)
}
