//go:generate mockgen -source=redis.go -destination=../../test/mock/external/redis_mock.go -package=mock
package external

import (
	"backend/core"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	RedisClient interface {
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		Del(ctx context.Context, keys ...string) error
		Exists(ctx context.Context, keys ...string) (int64, error)
		Close() error
	}

	redisClient struct {
		*redis.Client
	}
)

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *redisClient) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}

func (r *redisClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.Client.Exists(ctx, keys...).Result()
}

func (r *redisClient) Close() error {
	return r.Client.Close()
}

func NewRedisClient(cfg core.RedisConfig) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	return &redisClient{rdb}
}
