package record

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	DocumentNumber string          `gorm:"not null;unique;type:varchar(20)"`
	Balance        decimal.Decimal `gorm:"type:decimal(20,5);default:0.0"`
}

type Transaction struct {
	gorm.Model
	AccountID      uint `gorm:"foreignKey:AccountId;references:ID;not null;index;constraints:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OperationType  int
	Amount         decimal.Decimal `gorm:"type:decimal(20,5);not null"`
	EventTimestamp time.Time
}
