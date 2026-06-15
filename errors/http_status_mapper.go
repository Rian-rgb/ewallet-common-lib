package errors

import "net/http"

func (c Code) ToHTTPStatus() int {
	switch c {
	case ErrCodeUnknownError:
		return http.StatusInternalServerError

	case ErrCodeBadRequest, ErrCodeInvalidStatusTransition:
		return http.StatusBadRequest

	case ErrCodeUnauthorized, ErrCodeInvalidPassword:
		return http.StatusUnauthorized

	case ErrCodeForbidden:
		return http.StatusForbidden

	case ErrCodeNotFound, ErrCodeUserNotFound, ErrCodeTransactionNotFound, ErrCodeSessionNotFound:
		return http.StatusNotFound

	case ErrCodeInsufficientBalance:
		return http.StatusUnprocessableEntity // HTTP 422 cocok untuk bisnis logic yang gagal

	default:
		return http.StatusInternalServerError
	}
}
