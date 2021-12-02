package frlimit

import (
	"context"
)

type MemLimiter struct {
	ctx context.Context
	Limiter
}

var (
	defaultLimiter *MemLimiter = NewMemLimiter()
)

// L return default memory limiter
func L() Limiter {
	return defaultLimiter
}

func NewMemLimiter() *MemLimiter {
	return &MemLimiter{
		ctx: context.Background(),
	}
}
