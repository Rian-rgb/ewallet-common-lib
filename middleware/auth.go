package middleware

import (
	"github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"strings"
)

type TokenValidatorFunc func(tokenString string) (*security.ClaimToken, error)

func AuthMiddleware(validateToken TokenValidatorFunc, redisRepo redis.RedisRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			errCodeUnauthorized        = errors.ErrCodeUnauthorized
			errCodeInternalServerError = errors.ErrCodeInternalServerError
		)

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			logger.WithContext(ctx).Error("authorization header is empty")
			response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.UnauthorizedMessage)
			ctx.Abort()
			return
		}

		tokenString := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		claim, err := validateToken(tokenString)
		if err != nil {
			logger.WithContext(ctx).Error("failed to validate token: ", err)
			response.SendError(
				ctx,
				errCodeUnauthorized.ToHTTPStatus(),
				errCodeUnauthorized,
				response.InvalidTokenMessage,
			)
			ctx.Abort()
			return
		}

		userTokenKey := redis.UserTokenPrefix + tokenString
		exists, err := redisRepo.Exists(ctx, userTokenKey)
		if err != nil {
			logger.WithContext(ctx).Error("failed to get token from redis: ", err)
			response.SendError(
				ctx,
				errCodeInternalServerError.ToHTTPStatus(),
				errCodeInternalServerError,
				response.InternalServerErrorMessage,
			)
			ctx.Abort()
			return
		}

		if !exists {
			logger.WithContext(ctx).Error("token no exists in redis")
			response.SendError(
				ctx,
				errCodeUnauthorized.ToHTTPStatus(),
				errCodeUnauthorized,
				response.InvalidTokenMessage,
			)
			ctx.Abort()
		}

		tokenData := security.Token{
			UserID:   claim.UserID,
			Username: claim.Username,
			FullName: claim.FullName,
		}
		security.SetGinToken(ctx, tokenData)
		ctx.Next()
	}
}
