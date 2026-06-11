package errors

import "fmt"

type AppError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewAppError(code Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func ParseError(err error) (*AppError, bool) {
	if err == nil {
		return nil, false
	}
	appErr, ok := err.(*AppError)
	return appErr, ok
}
