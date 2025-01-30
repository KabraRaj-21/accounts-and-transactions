package handler

import (
	"transaction/internal/logger"
	"transaction/internal/server/http/utils"
	"transaction/internal/transaction/types/transaction_service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService transaction_service.TransactionService
}

func NewTransactionHandler(transactionService transaction_service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (th *TransactionHandler) RegisterTransaction(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	request := &transaction_service.RegisterTransactionRequest{}
	err := utils.UnmarshalJsonRequest(ginCtx, request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error unmarshalling request err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}
	err = utils.Validate(request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error validating request err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}

	res, err := th.transactionService.RegisterTransaction(ctx, request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error registring transaction err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}

	utils.SendSuccesResponse(ginCtx, "Success", res)
}
