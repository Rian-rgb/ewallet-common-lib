package redis

import "time"

type ExpiredDuration time.Duration

const (
	RefreshTokenDuration ExpiredDuration = ExpiredDuration(72 * time.Hour)
)
