package response

const (
	SuccessMessage             = "success"
	UnauthorizedMessage        = "authorization token is required"
	InvalidTokenMessage        = "invalid token"
	TokenExpiredMessage        = "your token has expired. Please login again"
	InternalServerErrorMessage = "an unexpected error occurred. Please try again later"
	InvalidRequestMessage      = "invalid request"
	InvalidJSONFormatMessage   = "invalid JSON format"
)
