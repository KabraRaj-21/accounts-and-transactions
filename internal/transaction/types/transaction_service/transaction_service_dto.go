package transaction_service

type RegisterTransactionRequest struct {
	AccountId     string  `json:"account_id" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
	OperationType int     `json:"operation_type_id" validate:"required"`
}
