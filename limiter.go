package frlimit

import (
	"github.com/smokecat/frlimit/adapter"
)

type Limiter struct {
	adapter *Adapter
}

var (
	defaultMemAdapter          = adapter.NewMemAdapter()
	defaultLimiter    *Limiter = NewLimiter(&defaultMemAdapter)
)

func NewLimiter(adapter *Adapter) *Limiter {
	return &Limiter{adapter: adapter}
}
