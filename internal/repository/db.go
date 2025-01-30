package repository

import (
	"accounts-and-transactions/internal/config"
	"accounts-and-transactions/internal/datasource/mysql"
	"accounts-and-transactions/internal/repository/types/record"
	"context"

	"gorm.io/gorm"
)

func NewAccountsSQLDB(ctx context.Context, config *config.DBConfig) (*gorm.DB, error) {
	conn, err := mysql.NewMySQLConnection(ctx, config.UserName, config.Password, config.DatabaseName, config.Host, config.Port)
	if err != nil {
		return nil, err
	}

	gormDB, err := mysql.NewGROMDB(ctx, conn)
	if err != nil {
		return nil, err
	}

	if config.AutoMigrateRequired {
		autoMigrate(gormDB)
	}
	return gormDB, nil
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(record.Account{})
	db.AutoMigrate(record.Transaction{})
}
