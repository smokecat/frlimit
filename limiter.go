package frlimit

import (
	"time"
)

type Limiter interface {
	// Check true if consume given amount successfully.
	Check(key string, amount, capacity int64, period time.Duration) (bool, error)

	// Count increase and return count.
	Count(key string) (int, error)

	// ResetCheck delete check key.
	ResetCheck(key string) (exists bool, err error)

	// ResetCount delete count key.
	ResetCount(key string) (exists bool, err error)
}
