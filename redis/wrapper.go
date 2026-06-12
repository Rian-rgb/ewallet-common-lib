package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Wrapper struct {
	core *redis.Client
}

func (w *Wrapper) Get(ctx context.Context, key string) (string, error) {
	return w.core.Get(ctx, key).Result()
}

func (w *Wrapper) Exists(ctx context.Context, key string) (bool, error) {
	n, err := w.core.Exists(ctx, key).Result()
	return n > 0, err
}

func (w *Wrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return w.core.Set(ctx, key, value, expiration).Err()
}

func (w *Wrapper) Delete(ctx context.Context, key string) error {
	return w.core.Del(ctx, key).Err()
}
