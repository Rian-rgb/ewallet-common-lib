package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewClient(
	addr string,
	password string,
	db int,
) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
