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

	case ErrCodeNotFound,
		ErrCodeUserNotFound,
		ErrCodeTransactionNotFound,
		ErrCodeSessionNotFound,
		ErrCodeWalletNotFound:

		return http.StatusNotFound

	case ErrCodeDuplicateReference:
		return http.StatusConflict

	case ErrCodeInsufficientBalance:
		return http.StatusUnprocessableEntity // HTTP 422 suitable for bussiness logic failed

	default:
		return http.StatusInternalServerError
	}
}
