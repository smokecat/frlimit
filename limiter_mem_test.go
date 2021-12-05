package frlimit

import (
	"context"
	"math"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestL(t *testing.T) {
	tests := []struct {
		name string
		want Limiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := L(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("L() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemLimiter_Check(t *testing.T) {
	type fields struct {
		ctx     context.Context
		bucket  map[string]*bucket
		store   *sync.Map
		Limiter Limiter
	}
	type args struct {
		key      string
		amount   int64
		capacity int64
		period   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"check 5/5", fields{ctx: context.Background(), bucket: map[string]*bucket{}, store: &sync.Map{}},
			args{"check 5/5", 1, 5, 5 * time.Second}, true, false},
		{"check 5/10", fields{ctx: context.Background(), bucket: map[string]*bucket{}, store: &sync.Map{}},
			args{"check 5/10", 1, 5, 10 * time.Second}, true, false},
		{"check 10/5", fields{ctx: context.Background(), bucket: map[string]*bucket{}, store: &sync.Map{}},
			args{"check 10/5", 1, 10, 5 * time.Second}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MemLimiter{
				ctx:     tt.fields.ctx,
				bucket:  tt.fields.bucket,
				store:   tt.fields.store,
				Limiter: tt.fields.Limiter,
			}

			testLimiterCheck(l, tt.args.key, tt.args.amount, tt.args.capacity, tt.args.period, t)
		})
	}
}

func TestNewMemLimiter(t *testing.T) {
	tests := []struct {
		name string
		want *MemLimiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemLimiter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemLimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testLimiterCheck(l Limiter, key string, amount, capacity int64, period time.Duration, t *testing.T) {
	var (
		total, allowed float64 = 0, 0
	)
	// 每0.1秒check一次，统计速率
	ticker := time.NewTicker(10 * time.Millisecond)

	_, _ = l.Check(key, capacity-amount, capacity, period)
	go func() {
		for {
			select {
			case <-ticker.C:
				b, err := l.Check(key, amount, capacity, period)
				if err != nil {
					t.Errorf("Check err: %v", err)
				}
				total++
				if b {
					allowed++
				}
			}
		}
	}()

	seconds := 10 * time.Second
	time.Sleep(seconds)
	ticker.Stop()

	gotRate := allowed / seconds.Seconds()
	expectRate := float64(capacity) / period.Seconds()
	// 统计
	t.Logf("Check key = %s amount = %d capacity = %d period = %d: total = %g allowed = %g gotRate = %g count/s", key, amount,
		capacity, period, total, allowed, gotRate)
	if math.Abs(gotRate-expectRate) > 1/seconds.Seconds() {
		t.Errorf("Check key = %s capacity = %d period = %d: got gotRate = %g expect gotRate = %g", key, capacity, period, gotRate,
			expectRate)
	}
}
