package frlimit

import "context"

type Adapter interface {
	// Get(ctx context.Context, key string) (int, error)
	// Set(ctx context.Context, key string, value int) error
	// Incr(ctx context.Context, key string) (int, error)
	// Del(ctx context.Context, key string) error
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)
	EvalSha(ctx context.Context, sha string, keys []string, args ...interface{}) (interface{}, error)
	EvalBool(ctx context.Context, script string, keys []string, args ...interface{}) (bool, error)
	EvalShaBool(ctx context.Context, sha string, keys []string, args ...interface{}) (bool, error)
	EvalString(ctx context.Context, script string, keys []string, args ...interface{}) (string, error)
	EvalShaString(ctx context.Context, sha string, keys []string, args ...interface{}) (string, error)
	EvalInt(ctx context.Context, script string, keys []string, args ...interface{}) (int, error)
	EvalShaInt(ctx context.Context, sha string, keys []string, args ...interface{}) (int, error)

	ScriptExists(ctx context.Context, shas ...string) ([]bool, error)
	ScriptFlush(ctx context.Context) error
	ScriptKill(ctx context.Context) error
	ScriptLoad(ctx context.Context, script string) (string, error)
}
