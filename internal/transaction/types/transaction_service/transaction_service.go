package transaction_service

import (
	"context"
	"transaction/internal/entity"
)

type TransactionService interface {
	RegisterTransaction(ctx context.Context, req *RegisterTransactionRequest) (*entity.Transaction, error)
}
