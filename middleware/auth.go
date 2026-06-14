package middleware

import (
	"github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type TokenValidatorFunc func(tokenString string) (*security.ClaimToken, error)

func AuthMiddleware(validateToken TokenValidatorFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		codeUnauthorized := errors.ErrUnauthorized

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.WithContext(c.Request.Context()).Error("authorization header is empty")
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), response.UnauthorizedMessage)
			c.Abort()
			return
		}

		tokenString := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		claim, err := validateToken(tokenString)
		if err != nil {
			logger.WithContext(c.Request.Context()).Error("failed to validate token: %v", err)
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), response.InvalidTokenMessage)
			c.Abort()
			return
		}

		expTime, err := claim.GetExpirationTime()
		if err != nil || time.Now().After(expTime.Time) {
			logger.WithContext(c).Error("token has expired, expired at: %v", claim.ExpiresAt)
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), response.TokenExpiredMessage)
			c.Abort()
			return
		}

		c.Set("token", claim)
		c.Next()
	}
}
