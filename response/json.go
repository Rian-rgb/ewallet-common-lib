package response

import (
	"github.com/gin-gonic/gin"
)

func SendSuccess(c *gin.Context, httpCode int, message string, data interface{}) {
	c.JSON(httpCode, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendSuccessWithMeta(c *gin.Context, httpCode int, message string, data interface{}, meta PaginationMeta) {
	c.JSON(httpCode, SuccessResponse{
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func SendError(c *gin.Context, httpCode int, errorCode string, message string) {
	c.JSON(httpCode, ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	})
}

func SendBadRequest(c *gin.Context, errorCode string, message string, validationErrors []ValidationErrorField) {
	c.JSON(400, BadRequestResponse{
		ErrorCode: errorCode,
		Message:   message,
		Errors:    validationErrors,
	})
}
