package jwt

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

type RueidisCache struct {
	client rueidis.Client
}

func NewRueidisCache(client rueidis.Client) *RueidisCache {
	return &RueidisCache{client: client}
}

func (c *RueidisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Do(ctx, c.client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *RueidisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).Ex(expiration).Build()).Error()
}

func (c *RueidisCache) Del(ctx context.Context, key string) error {
	return c.client.Do(ctx, c.client.B().Del().Key(key).Build()).Error()
}
