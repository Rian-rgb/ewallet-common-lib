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

func RefreshTokenMiddleware(
	validateToken TokenValidatorFunc,
	redisRepo redis.RedisRepository,
	secretKeyEncrypt string,
) gin.HandlerFunc {
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

		encryptRefreshToken := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			encryptRefreshToken = strings.TrimPrefix(authHeader, "Bearer ")
		}

		refreshToken, err := security.Decrypt(encryptRefreshToken, []byte(secretKeyEncrypt))

		claim, err := validateToken(string(refreshToken))
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

		refreshTokenKey := redis.RefreshTokenPrefix + claim.ID
		exists, err := redisRepo.Exists(ctx, refreshTokenKey)
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
			logger.WithContext(ctx).Error("refresh token no exists in redis")
			response.SendError(
				ctx,
				errCodeUnauthorized.ToHTTPStatus(),
				errCodeUnauthorized,
				response.InvalidTokenMessage,
			)
			ctx.Abort()
		}

		refreshTokenData := security.Token{
			UserID:   claim.UserID,
			Username: claim.Username,
			FullName: claim.FullName,
		}
		security.SetGinToken(ctx, refreshTokenData)
		ctx.Next()
	}
}
