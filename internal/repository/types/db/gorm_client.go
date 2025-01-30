package db

import (
	"database/sql"

	"gorm.io/gorm"
)

type GORMClient interface {
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
}
