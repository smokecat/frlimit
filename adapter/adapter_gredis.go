package adapter

import (
	"context"

	"github.com/gogf/gf/database/gredis"
)

type adapterGredis struct {
	redis *gredis.Redis
}

func NewAdapterGRedis(redis *gredis.Redis) *adapterGredis {
	return &adapterGredis{redis}
}

func (a *adapterGredis) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalSha(ctx context.Context, sha string, keys []string, args ...interface{}) (interface{}, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalBool(ctx context.Context, script string, keys []string, args ...interface{}) (bool, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalShaBool(ctx context.Context, sha string, keys []string, args ...interface{}) (bool, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalString(ctx context.Context, script string, keys []string, args ...interface{}) (string, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalShaString(ctx context.Context, sha string, keys []string, args ...interface{}) (string, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalInt(ctx context.Context, script string, keys []string, args ...interface{}) (int, error) {
	panic("implement me")
}

func (a *adapterGredis) EvalShaInt(ctx context.Context, sha string, keys []string, args ...interface{}) (int, error) {
	panic("implement me")
}

func (a *adapterGredis) ScriptExists(ctx context.Context, shas ...string) ([]bool, error) {
	panic("implement me")
}

func (a *adapterGredis) ScriptFlush(ctx context.Context) error {
	panic("implement me")
}

func (a *adapterGredis) ScriptKill(ctx context.Context) error {
	panic("implement me")
}

func (a *adapterGredis) ScriptLoad(ctx context.Context, script string) (string, error) {
	panic("implement me")
}

func (a *adapterGredis) Get(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (a *adapterGredis) Set(ctx context.Context, key string, value int) error {
	return nil
}

func (a *adapterGredis) Incr(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (a *adapterGredis) Del(ctx context.Context, key string) error {
	return nil
}
