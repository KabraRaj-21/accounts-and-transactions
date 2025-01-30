package handler

import (
	"accounts-and-transactions/internal/account/types/account_service"
	"accounts-and-transactions/internal/logger"
	"accounts-and-transactions/internal/server/http/utils"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService account_service.AccountService
}

func NewAccountHandler(accountService account_service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (ah *AccountHandler) CreateAccount(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	request := &account_service.CreateAccountRequest{}
	err := utils.UnmarshalJsonRequest(ginCtx, request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error unmarshalling request, err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}
	err = utils.Validate(request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error validating request, err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}

	res, err := ah.accountService.Create(ctx, request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error creating account err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}

	utils.SendSuccesResponse(ginCtx, "Success", res)
}

func (ah *AccountHandler) GetAccount(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	request := &account_service.GetAccountRequest{
		Id: ginCtx.Param("accountId"),
	}
	err := utils.Validate(request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error validating request err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}

	res, err := ah.accountService.Get(ctx, request)
	if err != nil {
		logger.WithContext(ctx).Errorf("error getting account err:%s", err.Error())
		utils.SendErrorResponse(ginCtx, err)
		return
	}

	utils.SendSuccesResponse(ginCtx, "Success", res)
}
