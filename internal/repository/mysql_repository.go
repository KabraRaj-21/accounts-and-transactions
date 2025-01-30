package repository

import (
	"accounts-and-transactions/internal/entity"
	"accounts-and-transactions/internal/errors/tserror"
	"accounts-and-transactions/internal/repository/mapper"
	"accounts-and-transactions/internal/repository/types/db"
	"accounts-and-transactions/internal/repository/types/record"
	"context"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db db.GORMClient
}

func NewMySQLRepository(db db.GORMClient) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) CreateAccount(ctx context.Context, req *entity.Account) (*entity.Account, error) {
	accountRecord := mapper.MapAccountEntityToRecord(req)

	res := r.db.Create(accountRecord)
	if res.Error != nil {
		return nil, tserror.MapGormError(res.Error)
	} else if res.RowsAffected == 0 {
		return nil, tserror.New(tserror.ErrorType_INVALID_REQUEST, "duplicate request")
	}
	return mapper.MapAccountRecordToEntity(accountRecord), nil
}

func (r *MySQLRepository) GetAccountById(ctx context.Context, id string) (*entity.Account, error) {
	var accountRecord record.Account
	res := r.db.First(&accountRecord, id)
	if res.Error != nil {
		return nil, tserror.MapGormError(res.Error)
	}
	return mapper.MapAccountRecordToEntity(&accountRecord), nil
}

func (r *MySQLRepository) RegisterTransaction(ctx context.Context, account *entity.Account, transaction *entity.Transaction) (*entity.Transaction, error) {
	id, err := mapper.GetAccountRecordIdFromEntity(account)
	if err != nil {
		return nil, err
	}
	transactionRecord, err := mapper.MapTransactionEntityToRecord(transaction)
	if err != nil {
		return nil, err
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&record.Account{}).
			Where("id = ?", id).
			Update("balance", account.Balance).
			Error; err != nil {
			return err
		}

		if err := tx.Create(&transactionRecord).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, tserror.MapGormError(err)
	}
	return mapper.MapTransactionRecordToEntity(transactionRecord), nil
}
