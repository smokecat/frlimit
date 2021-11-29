package adapter

import (
	"context"

	"github.com/gogf/gf/database/gredis"

	"github.com/smokecat/frlimit"
)

type adapterGRedis struct {
	redis *gredis.Redis
}

func NewAdapterGRedis(redis *gredis.Redis) frlimit.Adapter {
	return &adapterGRedis{redis}
}

func (ma adapterGRedis) Get(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma adapterGRedis) Set(ctx context.Context, key string, value int) error {
	return nil
}

func (ma adapterGRedis) Incr(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma adapterGRedis) Del(ctx context.Context, key string) error {
	return nil
}
