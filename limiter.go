package frlimit

import (
	"github.com/smokecat/frlimit/adapter"
)

type Limiter struct {
	adapter *Adapter
}

var (
	defaultMemAdapter          = adapter.NewAdapterRedis()
	defaultLimiter    *Limiter = New(&defaultMemAdapter)
)

func New(adapter *Adapter) *Limiter {
	return &Limiter{adapter: adapter}
}
