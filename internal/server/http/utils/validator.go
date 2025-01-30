package utils

import (
	"transaction/internal/errors/tserror"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func Validate(input interface{}) error {
	if err := v.Struct(input); err != nil {
		return tserror.Wrap(tserror.ErrorType_INVALID_REQUEST, "invalid request", err)
	}
	return nil
}
