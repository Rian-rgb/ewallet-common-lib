package errors

import "net/http"

func (c Code) ToHTTPStatus() int {
	switch c {
	case ErrInternalServerError:
		return http.StatusInternalServerError

	case ErrBadRequest, ErrInvalidStatusTransition:
		return http.StatusBadRequest

	case ErrUnauthorized, ErrInvalidPassword:
		return http.StatusUnauthorized

	case ErrForbidden:
		return http.StatusForbidden

	case ErrNotFound, ErrUserNotFound, ErrTransactionNotFound, ErrSessionNotFound:
		return http.StatusNotFound

	case ErrInsufficientBalance:
		return http.StatusUnprocessableEntity // HTTP 422 cocok untuk bisnis logic yang gagal

	default:
		return http.StatusInternalServerError
	}
}
