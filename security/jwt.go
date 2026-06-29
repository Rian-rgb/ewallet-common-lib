package security

import (
	"fmt"
	"github.com/Rian-rgb/ewallet-common-lib/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ClaimToken struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
	issuer        string
}

func NewJWTManager(secret string, issuer string) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secret),
		issuer:    issuer,
	}
}

func (m *JWTManager) GenerateToken(
	userID int,
	username string,
	fullName string,
	expiration ExpiredDuration,
) (token string, jti string, err error) {

	now := time.Now()
	jti = utils.GenerateUUID()

	claimToken := ClaimToken{
		UserID:   userID,
		Username: username,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expiration))),
			ID:        jti,
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)
	token, err = generateToken.SignedString(m.secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, jti, nil
}

func (m *JWTManager) ValidateToken(tokenString string) (*ClaimToken, error) {

	claimToken := &ClaimToken{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, claimToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to validate method jwt: %v", t.Header["alg"])
		}
		return m.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token jwt: %w", err)
	}

	if !jwtToken.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claimToken, nil
}
