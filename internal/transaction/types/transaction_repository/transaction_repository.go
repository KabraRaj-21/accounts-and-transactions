package transaction_repository

import (
	"accounts-and-transactions/internal/entity"
	"context"
)

type TransactionRepository interface {
	RegisterTransaction(ctx context.Context, account *entity.Account, transaction *entity.Transaction) (*entity.Transaction, error)
}
