package tserror

import "fmt"

type ErrorType uint

const (
	ErrorType_UNKNOWN_FAILURE ErrorType = iota
	ErrorType_INVALID_REQUEST
	ErrorType_OPERATION_FAILURE
	ErrorType_NOT_FOUND
	ErrorType_UNAUTHORIZED_OPERATION
)

var tsErrorNameMap = map[ErrorType]string{
	ErrorType_UNKNOWN_FAILURE:        "UNKNOWN_FAILURE",
	ErrorType_INVALID_REQUEST:        "INVALID_REQUEST",
	ErrorType_OPERATION_FAILURE:      "OPEARTION_FAILURE",
	ErrorType_NOT_FOUND:              "NOT_FOUND",
	ErrorType_UNAUTHORIZED_OPERATION: "UNAUTHORIZED_OPERATION",
}

func (t ErrorType) String() string {
	return tsErrorNameMap[t]
}

type tsError struct {
	ErrorType     ErrorType
	Message       string
	OriginalError error
}

func (e *tsError) Error() string {
	if e.OriginalError != nil {
		return fmt.Sprintf("ErrorType:%s, Issue:%s Cause:%s", e.ErrorType.String(), e.Message, e.OriginalError.Error())
	}
	return fmt.Sprintf("ErrorType:%s, Issue:%s", e.ErrorType.String(), e.Message)
}

func New(t ErrorType, message string) error {
	return &tsError{
		ErrorType:     t,
		Message:       message,
		OriginalError: nil,
	}
}

func Wrap(t ErrorType, message string, originalError error) error {
	return &tsError{
		ErrorType:     t,
		Message:       message,
		OriginalError: originalError,
	}
}
