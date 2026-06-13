package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	Core *redis.Client
}

func (w *RedisRepository) Exists(ctx context.Context, key string) (bool, error) {
	n, err := w.Core.Exists(ctx, key).Result()
	return n > 0, err
}

func (w *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration ExpiredDuration) error {
	return w.Core.Set(ctx, key, value, time.Duration(expiration)).Err()
}

func (w *RedisRepository) Delete(ctx context.Context, key string) error {
	return w.Core.Del(ctx, key).Err()
}
