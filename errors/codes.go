package errors

type Code string

const (

	// ---- GENERAL ERRORS ----
	ErrCodeUnknownError        Code = "UNKNOWN_ERROR"
	ErrCodeInternalServerError Code = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest          Code = "BAD_REQUEST"
	ErrCodeUnauthorized        Code = "UNAUTHORIZED"
	ErrCodeForbidden           Code = "FORBIDDEN"
	ErrCodeNotFound            Code = "NOT_FOUND"

	// ---- AUTH & USER ERRORS ----
	ErrCodeUserNotFound    Code = "USER_NOT_FOUND"
	ErrCodeSessionNotFound Code = "SESSION_NOT_FOUND"
	ErrCodeInvalidPassword Code = "INVALID_PASSWORD"

	// ---- WALLET & TRANSACTION ERRORS ----
	ErrCodeWalletNotFound          Code = "WALLET_NOT_FOUND"
	ErrCodeInvalidStatusTransition Code = "INVALID_STATUS_TRANSITION"
	ErrCodeDuplicateReference      Code = "DUPLICATE_REFERENCE"
	ErrCodeTransactionNotFound     Code = "TRANSACTION_NOT_FOUND"
	ErrCodeInsufficientBalance     Code = "INSUFFICIENT_BALANCE"
)
