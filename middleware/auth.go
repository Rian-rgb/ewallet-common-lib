package middleware

import (
	"github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type TokenValidatorFunc func(tokenString string) (*CustomClaims, error)

type CustomClaims struct {
	UserID    string
	Email     string
	ExpiresAt time.Time
}

func AuthMiddleware(validateToken TokenValidatorFunc, redisCl redis.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			codeUnauthorized    = errors.ErrUnauthorized
			codeInternalError   = errors.ErrInternalError
			codeSessionNotFound = errors.ErrSessionNotFound
		)

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

		if time.Now().After(claim.ExpiresAt) {
			logger.WithContext(c.Request.Context()).Error("token has expired, expired at: %v", claim.ExpiresAt)
			response.SendError(c, codeUnauthorized.ToHTTPStatus(), string(codeUnauthorized), "Your token has expired. Please login again.")
			c.Abort()
			return
		}

		sessionKey := redis.SessionPrefix + tokenString
		isSessioned, err := redisCl.Exists(c, sessionKey)
		if err != nil {
			logger.WithContext(c).Error("redis error when checking token session: %v", err)
			response.SendError(c, codeInternalError.ToHTTPStatus(), string(codeInternalError), "An unexpected error occurred. Please try again later.")
			c.Abort()
			return
		}

		if !isSessioned {
			logger.WithContext(c).Error("failed to exist user session")
			response.SendError(c, codeSessionNotFound.ToHTTPStatus(), string(codeSessionNotFound), "User session not found.")
			c.Abort()
			return
		}

		c.Set("token", claim)
		c.Next()
	}
}
