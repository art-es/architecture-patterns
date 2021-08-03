// Package redis is an adapter between github.com/go-redis/redis and internal domain layer
package redis

import (
	"context"
	"time"

	"github.com/art-es/architecture-patterns/layered-pattern/util/json"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	return c.Client.Get(ctx, key).Bytes()
}

func (c *Client) SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	valueInJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Client.Set(ctx, key, valueInJSON, expiration).Err()
}
