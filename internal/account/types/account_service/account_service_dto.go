package account_service

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

type GetAccountRequest struct {
	Id string `validate:"required"`
}
