package security

import (
	"github.com/gin-gonic/gin"
)

const GinTokenKey = "TokenClaims"

func SetGinToken(ctx *gin.Context, token Token) {
	ctx.Set(GinTokenKey, token)
}

func GetGinToken(ctx *gin.Context) (Token, bool) {
	val, exists := ctx.Get(GinTokenKey)
	if !exists {
		return Token{}, false
	}

	token, ok := val.(Token)
	return token, ok
}
