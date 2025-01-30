package account_info_service

import (
	"context"
	"transaction/internal/account/types/account_service"
	"transaction/internal/entity"
)

type AccountInfoService interface {
	Get(ctx context.Context, req *account_service.GetAccountRequest) (*entity.Account, error)
}
