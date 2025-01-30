package transaction_service

import (
	"accounts-and-transactions/internal/entity"
	"context"
)

type TransactionService interface {
	RegisterTransaction(ctx context.Context, req *RegisterTransactionRequest) (*entity.Transaction, error)
}
