package redis

import (
	"context"
)

type Repository interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value interface{}, expiration ExpiredDuration) error
	Delete(ctx context.Context, key string) error
}
