package mysql

import (
	"accounts-and-transactions/internal/logger"
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func NewGROMDB(ctx context.Context, db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		logger.WithContext(ctx).Errorf("error creating gorm db, err: %v", err)
		return nil, err
	}
	return gormDB, err
}

func NewMySQLConnection(ctx context.Context, userName, password, database, host, port string) (*sql.DB, error) {
	dbConnectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, host, port, database)

	db, err := sql.Open("mysql", dbConnectionURL)
	if err != nil {
		logger.WithContext(ctx).Errorf("error connecting to mysql database, err: %v", err)
		return nil, err
	}
	return db, nil
}
