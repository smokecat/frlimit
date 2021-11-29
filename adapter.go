package frlimit

import (
	"context"
)

type Adapter interface {
	Get(ctx context.Context, key string) (int, error)
	Set(ctx context.Context, key string, value int) error
	Incr(ctx context.Context, key string) (int, error)
	Del(ctx context.Context, key string) error
}
