package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	core *redis.Client
}

func (w *RedisRepository) Exists(ctx context.Context, key string) (bool, error) {
	n, err := w.core.Exists(ctx, key).Result()
	return n > 0, err
}

func (w *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return w.core.Set(ctx, key, value, expiration).Err()
}

func (w *RedisRepository) Delete(ctx context.Context, key string) error {
	return w.core.Del(ctx, key).Err()
}
