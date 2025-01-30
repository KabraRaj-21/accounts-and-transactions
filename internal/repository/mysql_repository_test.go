//go:build unit

package repository

import (
	"accounts-and-transactions/internal/entity"
	"accounts-and-transactions/internal/repository/mocks"
	"accounts-and-transactions/internal/repository/types/record"
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestMySQLRepository_Create(t *testing.T) {
	// Test cases
	tests := map[string]struct {
		input            *entity.Account
		dbErr            error
		rowsAffected     int64
		expectedResponse *entity.Account
		isErrorExpected  bool
	}{
		"DB Error": {
			input:           &entity.Account{DocumentNumber: "1234", Balance: decimal.NewFromFloat(1000)},
			dbErr:           errors.New("db error"),
			isErrorExpected: true,
		},
		"No rows affected": {
			input:           &entity.Account{Balance: decimal.NewFromFloat(1000)},
			dbErr:           nil,
			rowsAffected:    0,
			isErrorExpected: true,
		},
		"Success": {
			input:            &entity.Account{Balance: decimal.NewFromFloat(1000.98)},
			dbErr:            nil,
			rowsAffected:     1,
			expectedResponse: &entity.Account{Balance: decimal.NewFromFloat(1000.98)},
			isErrorExpected:  false,
		},
	}

	for tcName, tt := range tests {
		t.Run(tcName, func(t *testing.T) {
			mockDB := mocks.NewGORMClient(t)
			repo := NewMySQLRepository(mockDB)

			// Mock DB behavior
			mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: tt.dbErr, RowsAffected: tt.rowsAffected})

			// Act
			result, err := repo.CreateAccount(context.Background(), tt.input)

			// Assert
			if tt.isErrorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, tt.expectedResponse.Balance.Equal(result.Balance))
				assert.Equal(t, tt.expectedResponse.DocumentNumber, result.DocumentNumber)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

func TestMySQLRepository_GetAccountById(t *testing.T) {
	// Test cases
	tests := map[string]struct {
		input            string
		dbErr            error
		dbRecord         *record.Account
		expectedResponse *entity.Account
		isErrorExpected  bool
	}{
		"DB Error": {
			input:           "123",
			dbErr:           gorm.ErrRecordNotFound,
			isErrorExpected: true,
		},
		"Success": {
			input:            "123",
			dbErr:            nil,
			dbRecord:         &record.Account{DocumentNumber: "456", Balance: decimal.NewFromFloat(1000.98)},
			expectedResponse: &entity.Account{Id: "123", DocumentNumber: "456", Balance: decimal.NewFromFloat(1000.98)},
			isErrorExpected:  false,
		},
	}

	for tcName, tt := range tests {
		t.Run(tcName, func(t *testing.T) {
			mockDB := mocks.NewGORMClient(t)
			repo := &MySQLRepository{db: mockDB}

			// Mock DB behavior
			mockDB.On("First", mock.Anything, tt.input).Run(func(args mock.Arguments) {
				if tt.dbRecord == nil {
					return
				}
				accountRecord := args.Get(0).(*record.Account)
				accountRecord.ID = 123
				accountRecord.Balance = tt.dbRecord.Balance
				accountRecord.DocumentNumber = tt.dbRecord.DocumentNumber
			}).Return(&gorm.DB{Error: tt.dbErr})

			// Act
			result, err := repo.GetAccountById(context.Background(), tt.input)

			// Assert
			if tt.isErrorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, result)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

func TestMySQLRepository_RegisterTransaction(t *testing.T) {
	validAccountEntity := &entity.Account{Id: "123", DocumentNumber: "456", Balance: decimal.NewFromFloat(100.34)}
	validTransactionEntity := &entity.Transaction{AccountId: "123", OperationType: entity.OperationType_CREDIT_VOUCHER, Amount: decimal.NewFromFloat(44.78)}
	// Test cases
	tests := map[string]struct {
		inputAccount         *entity.Account
		inputTransaction     *entity.Transaction
		isDBCallExpected     bool
		dbErr                error
		newTransactionRecord *record.Transaction
		expectedResponse     *entity.Transaction
		isErrorExpected      bool
	}{
		"invalid account": {
			inputAccount:     &entity.Account{Id: "abc", DocumentNumber: "123"},
			isDBCallExpected: false,
			isErrorExpected:  true,
		},
		"invalid transaction": {
			inputAccount:     validAccountEntity,
			inputTransaction: &entity.Transaction{AccountId: "abc"},
			isDBCallExpected: false,
			isErrorExpected:  true,
		},
		"transaction failure": {
			inputAccount:     validAccountEntity,
			inputTransaction: validTransactionEntity,
			isDBCallExpected: true,
			dbErr:            gorm.ErrInvalidTransaction,
			isErrorExpected:  true,
		},
		"success": {
			inputAccount:     validAccountEntity,
			inputTransaction: validTransactionEntity,
			isDBCallExpected: true,
			dbErr:            nil,
			expectedResponse: validTransactionEntity,
			isErrorExpected:  false,
		},
	}

	for tcName, tt := range tests {
		t.Run(tcName, func(t *testing.T) {
			mockDB := mocks.NewGORMClient(t)
			repo := &MySQLRepository{db: mockDB}

			if tt.isDBCallExpected {
				mockDB.On("Transaction", mock.Anything).Return(tt.dbErr)
			}

			// Act
			result, err := repo.RegisterTransaction(context.Background(), tt.inputAccount, tt.inputTransaction)

			// Assert
			if tt.isErrorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse.AccountId, result.AccountId)
				assert.True(t, tt.expectedResponse.Amount.Equal(result.Amount))
				assert.Equal(t, tt.expectedResponse.OperationType, result.OperationType)
				assert.Equal(t, tt.expectedResponse.Timestamp, result.Timestamp)
			}

			mockDB.AssertExpectations(t)
		})
	}
}
