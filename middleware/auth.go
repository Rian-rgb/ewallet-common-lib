package middleware

import (
	"github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type TokenValidatorFunc func(tokenString string) (*security.ClaimToken, error)

func AuthMiddleware(validateToken TokenValidatorFunc, redisRepo redis.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		codeUnauthorized := errors.ErrUnauthorized

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.WithContext(c.Request.Context()).Error("authorization header is empty")
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), "Authorization token is required.")
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
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), "Invalid token.")
			c.Abort()
			return
		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			logger.WithContext(c.Request.Context()).Error("token has expired, expired at: %v", claim.ExpiresAt)
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), "Your token has expired. Please login again.")
			c.Abort()
			return
		}

		c.Set("token", claim)
		c.Next()
	}
}
