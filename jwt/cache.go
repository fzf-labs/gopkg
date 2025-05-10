package jwt

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

type RedisCache struct {
	redis *redis.Client
}

func NewRedisCache(redis *redis.Client) *RedisCache {
	return &RedisCache{redis: redis}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.redis.Set(ctx, key, value, expiration).Err()
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
	return c.redis.Del(ctx, key).Err()
}
