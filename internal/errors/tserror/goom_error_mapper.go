package tserror

import (
	"errors"

	"gorm.io/gorm"
)

func MapGormError(e error) error {
	switch {
	case errors.Is(e, gorm.ErrRecordNotFound):
		return Wrap(ErrorType_NOT_FOUND, "the requested resource could not be found", e)
	case errors.Is(e, gorm.ErrDuplicatedKey):
		return Wrap(ErrorType_INVALID_REQUEST, "duplicate entry, resource already exists", e)
	case errors.Is(e, gorm.ErrCheckConstraintViolated):
		fallthrough
	case errors.Is(e, gorm.ErrForeignKeyViolated):
		return Wrap(ErrorType_INVALID_REQUEST, "operation not allowed", e)
	case errors.Is(e, gorm.ErrInvalidValue):
		fallthrough
	case errors.Is(e, gorm.ErrInvalidData):
		fallthrough
	case errors.Is(e, gorm.ErrPrimaryKeyRequired):
		return Wrap(ErrorType_INVALID_REQUEST, "invalid request", e)
	case errors.Is(e, gorm.ErrRegistered):
		return Wrap(ErrorType_INVALID_REQUEST, "resource is already registered", e)

	default:
		return Wrap(ErrorType_UNKNOWN_FAILURE, "unknown failure", e)
	}
}
