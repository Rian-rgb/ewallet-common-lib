package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CreateBlindIndex(plaintext string, blindIndexKey []byte) string {
	h := hmac.New(sha256.New, blindIndexKey)
	h.Write([]byte(plaintext))
	return hex.EncodeToString(h.Sum(nil))
}
