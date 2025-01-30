package utils

import (
	"net/http"
	"transaction/internal/errors/tserror"
	"transaction/internal/server/http/types"

	"github.com/gin-gonic/gin"
)

func SendSuccesResponse(ginCtx *gin.Context, message string, data interface{}) {
	if len(message) == 0 {
		message = http.StatusText(http.StatusOK)
	}
	ginCtx.JSON(http.StatusOK, createAPIResponse("OK", message, data))
}

func SendErrorResponse(ginCtx *gin.Context, err error) {
	statusCode := tserror.GetHttpStatusCodeFromError(err)
	ginCtx.JSON(statusCode, createAPIResponse("ERROR", http.StatusText(statusCode), nil))
}

func createAPIResponse(status, message string, data interface{}) types.APIResponse {
	apiResponse := types.APIResponse{
		Status:  status,
		Message: message,
	}
	if data != nil {
		apiResponse.Data = data
	}
	return apiResponse
}
