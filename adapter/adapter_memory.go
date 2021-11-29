package adapter

import (
	"context"

	"github.com/smokecat/frlimit"
)

type adapterMemory struct {
	ctx context.Context
}

func NewAdapterMemory() frlimit.Adapter {
	return &adapterMemory{context.Background()}
}

func (ma adapterMemory) Get(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma adapterMemory) Set(ctx context.Context, key string, value int) error {
	return nil
}

func (ma adapterMemory) Incr(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (ma adapterMemory) Del(ctx context.Context, key string) error {
	return nil
}
