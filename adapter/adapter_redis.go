package adapter

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type adapterRedis struct {
	redis *redis.Client
}

func NewAdapterRedis() *adapterRedis {
	return &adapterRedis{redis.NewClient(&redis.Options{})}
}

func (a *adapterRedis) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalSha(ctx context.Context, sha string, keys []string, args ...interface{}) (interface{}, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalBool(ctx context.Context, script string, keys []string, args ...interface{}) (bool, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalShaBool(ctx context.Context, sha string, keys []string, args ...interface{}) (bool, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalString(ctx context.Context, script string, keys []string, args ...interface{}) (string, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalShaString(ctx context.Context, sha string, keys []string, args ...interface{}) (string, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalInt(ctx context.Context, script string, keys []string, args ...interface{}) (int, error) {
	panic("implement me")
}

func (a *adapterRedis) EvalShaInt(ctx context.Context, sha string, keys []string, args ...interface{}) (int, error) {
	panic("implement me")
}

func (a *adapterRedis) ScriptExists(ctx context.Context, shas ...string) ([]bool, error) {
	panic("implement me")
}

func (a *adapterRedis) ScriptFlush(ctx context.Context) error {
	panic("implement me")
}

func (a *adapterRedis) ScriptKill(ctx context.Context) error {
	panic("implement me")
}

func (a *adapterRedis) ScriptLoad(ctx context.Context, script string) (string, error) {
	panic("implement me")
}

func (ma *adapterRedis) Get(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma *adapterRedis) Set(ctx context.Context, key string, value int) error {
	return nil
}

func (ma *adapterRedis) Incr(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma *adapterRedis) Del(ctx context.Context, key string) error {
	return nil
}
