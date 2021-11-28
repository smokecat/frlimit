package adapter

import (
	"context"

	"github.com/smokecat/frlimit"
)

type MemAdapter struct {
	ctx context.Context
}

func NewMemAdapter() frlimit.Adapter {
	return &MemAdapter{context.Background()}
}

func (ma MemAdapter) Get(key string) (int, error) {
	return 0, nil
}

func (ma MemAdapter) Set(key string, value int) error {
	return nil
}

func (ma MemAdapter) Incr(key string) (int, error) {
	return 0, nil
}

func (ma MemAdapter) Del(key string) error {
	return nil
}
