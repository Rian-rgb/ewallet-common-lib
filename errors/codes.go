package errors

type Code string

const (

	// ---- GENERAL ERRORS ----
	ErrInternalServerError Code = "INTERNAL_SERVER_ERROR"
	ErrBadRequest          Code = "BAD_REQUEST"
	ErrUnauthorized        Code = "UNAUTHORIZED"
	ErrForbidden           Code = "FORBIDDEN"
	ErrNotFound            Code = "NOT_FOUND"

	// ---- AUTH & USER ERRORS ----
	ErrUserNotFound    Code = "USER_NOT_FOUND"
	ErrSessionNotFound Code = "SESSION_NOT_FOUND"
	ErrInvalidPassword Code = "INVALID_PASSWORD"

	// ---- WALLET & TRANSACTION ERRORS ----
	ErrInvalidStatusTransition Code = "INVALID_STATUS_TRANSITION"
	ErrTransactionNotFound     Code = "TRANSACTION_NOT_FOUND"
	ErrInsufficientBalance     Code = "INSUFFICIENT_BALANCE"
)
