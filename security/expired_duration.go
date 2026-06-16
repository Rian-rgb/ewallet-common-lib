package security

import "time"

type ExpiredDuration time.Duration

const (
	UserTokenDuration    ExpiredDuration = ExpiredDuration(5 * time.Minute)
	RefreshTokenDuration ExpiredDuration = ExpiredDuration(72 * time.Hour)
)
