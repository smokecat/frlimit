package adapter

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/smokecat/frlimit"
)

type adapterRedis struct {
	redis *redis.Client
}

func NewAdapterRedis() frlimit.Adapter {
	return &adapterRedis{redis.NewClient(&redis.Options{})}
}

func (ma adapterRedis) Get(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma adapterRedis) Set(ctx context.Context, key string, value int) error {
	return nil
}

func (ma adapterRedis) Incr(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma adapterRedis) Del(ctx context.Context, key string) error {
	return nil
}
