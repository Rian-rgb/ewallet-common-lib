package context

import (
	"github.com/gin-gonic/gin"
)

const GinTokenKey = "TokenClaims"

func SetGinToken(c *gin.Context, token string) {
	c.Set(GinTokenKey, token)
}

func GetGinToken(c *gin.Context) (Token, bool) {
	val, exists := c.Get(GinTokenKey)
	if !exists {
		return Token{}, false
	}

	token, ok := val.(Token)
	return token, ok
}
