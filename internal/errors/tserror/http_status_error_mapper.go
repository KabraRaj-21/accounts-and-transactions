package tserror

import "net/http"

func GetHttpStatusCodeFromError(err error) int {
	if e, ok := err.(*tsError); ok {
		switch e.ErrorType {
		case ErrorType_INVALID_REQUEST:
			return http.StatusBadRequest
		case ErrorType_NOT_FOUND:
			return http.StatusNotFound
		case ErrorType_OPERATION_FAILURE:
			return http.StatusInternalServerError
		case ErrorType_UNAUTHORIZED_OPERATION:
			return http.StatusUnauthorized
		case ErrorType_UNKNOWN_FAILURE:
			fallthrough
		default:
			return http.StatusInternalServerError
		}
	}
	return http.StatusInternalServerError
}
