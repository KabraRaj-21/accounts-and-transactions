package validator_service

import "transaction/internal/entity"

type ActionType uint

const (
	ActionType_Unknown ActionType = iota
	ActionType_TransactionRegistration
)

type ValidationRequest struct {
	ActionType  ActionType
	Account     *entity.Account
	Transaction *entity.Transaction
}
