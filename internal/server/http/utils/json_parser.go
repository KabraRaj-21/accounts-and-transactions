package utils

import (
	"transaction/internal/errors/tserror"

	"github.com/gin-gonic/gin"
)

func UnmarshalJsonRequest(ginCtx *gin.Context, body interface{}) error {
	if err := ginCtx.ShouldBindJSON(body); err != nil {
		return tserror.Wrap(tserror.ErrorType_INVALID_REQUEST, "invalid request", err)
	}
	return nil
}
