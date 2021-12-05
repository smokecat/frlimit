package frlimit

import (
	"context"
	"errors"
	"sync"
	"time"
)

type bucket struct {
	capacity      float64
	period        float64
	leftCapacity  float64
	lastTimestamp int64
	mu            *sync.Mutex
}

type MemLimiter struct {
	ctx    context.Context
	bucket map[string]*bucket
	store  *sync.Map
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

// Check true if consume given amount successfully.
//
// Param period should use ms.
func (l *MemLimiter) Check(key string, amount, capacity int64, period time.Duration) (bool, error) {
	if amount > capacity {
		return false, errors.New("amount can not be greater than capacity")
	}

	var (
		famo, fcap, fperms         = float64(amount), float64(capacity), float64(period.Milliseconds())
		rate               float64 = fcap / fperms
		nowTimestamp       int64   = time.Now().UnixMilli()
	)

	// make bucket
	b, ok := l.bucket[key]
	if !ok {
		b = &bucket{
			capacity:      fcap,
			period:        fperms,
			leftCapacity:  fcap,
			lastTimestamp: 0,
			mu:            &sync.Mutex{},
		}
		l.bucket[key] = b
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	b.capacity++

	b.capacity, b.period = fcap, fperms

	// calculate the remaining capacity for this time
	leftCapacity, ok := l.calLeftCapacity(b, famo, fcap, fperms, rate, nowTimestamp)
	if !ok {
		return false, nil
	}

	// consume amount
	b.leftCapacity, b.lastTimestamp = leftCapacity-famo, nowTimestamp
	// l.bucket[key] = b
	return true, nil
}

// Count increase and return count.
func (l *MemLimiter) Count(key string) (int, error) {
	return 0, nil
}

// ResetCheck delete check key.
func (l *MemLimiter) ResetCheck(key string) (exists bool, err error) {
	return false, nil
}

// ResetCount delete count key.
func (l *MemLimiter) ResetCount(key string) (exists bool, err error) {
	return false, nil
}

func (l *MemLimiter) calLeftCapacity(b *bucket, amount, capacity, periodMs, rate float64, nowTimestamp int64) (float64,
	bool) {
	var leftCapacity = b.leftCapacity + float64(nowTimestamp-b.lastTimestamp)*rate

	if leftCapacity > capacity {
		leftCapacity = capacity
	}

	return leftCapacity, leftCapacity >= amount
}
