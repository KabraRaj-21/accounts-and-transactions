package transaction_repository

import (
	"context"
	"transaction/internal/entity"
)

type TransactionRepository interface {
	RegisterTransaction(ctx context.Context, account *entity.Account, transaction *entity.Transaction) (*entity.Transaction, error)
}
