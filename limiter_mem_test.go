package frlimit

import (
	"context"
	"reflect"
	"sync"
	"testing"
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
		bucket  map[string]bucket
		store   sync.Map
		Limiter Limiter
	}
	type args struct {
		key      string
		amount   int64
		capacity int64
		period   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"check 5/5", fields{ctx: context.Background(), bucket: map[string]bucket{}, store: sync.Map{}},
			args{"check 5/5", 1, 5, 5}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MemLimiter{
				ctx:     tt.fields.ctx,
				bucket:  tt.fields.bucket,
				store:   tt.fields.store,
				Limiter: tt.fields.Limiter,
			}
			got, err := l.Check(tt.args.key, tt.args.amount, tt.args.capacity, tt.args.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemLimiter_Count(t *testing.T) {
	type fields struct {
		ctx     context.Context
		bucket  map[string]bucket
		store   sync.Map
		Limiter Limiter
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MemLimiter{
				ctx:     tt.fields.ctx,
				bucket:  tt.fields.bucket,
				store:   tt.fields.store,
				Limiter: tt.fields.Limiter,
			}
			got, err := l.Count(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemLimiter_ResetCheck(t *testing.T) {
	type fields struct {
		ctx     context.Context
		bucket  map[string]bucket
		store   sync.Map
		Limiter Limiter
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantExists bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MemLimiter{
				ctx:     tt.fields.ctx,
				bucket:  tt.fields.bucket,
				store:   tt.fields.store,
				Limiter: tt.fields.Limiter,
			}
			gotExists, err := l.ResetCheck(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExists != tt.wantExists {
				t.Errorf("ResetCheck() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestMemLimiter_ResetCount(t *testing.T) {
	type fields struct {
		ctx     context.Context
		bucket  map[string]bucket
		store   sync.Map
		Limiter Limiter
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantExists bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MemLimiter{
				ctx:     tt.fields.ctx,
				bucket:  tt.fields.bucket,
				store:   tt.fields.store,
				Limiter: tt.fields.Limiter,
			}
			gotExists, err := l.ResetCount(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExists != tt.wantExists {
				t.Errorf("ResetCount() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestMemLimiter_calLeftCapacity(t *testing.T) {
	type fields struct {
		ctx     context.Context
		bucket  map[string]bucket
		store   sync.Map
		Limiter Limiter
	}
	type args struct {
		b            bucket
		amount       float64
		capacity     float64
		period       float64
		rate         float64
		nowTimestamp int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MemLimiter{
				ctx:     tt.fields.ctx,
				bucket:  tt.fields.bucket,
				store:   tt.fields.store,
				Limiter: tt.fields.Limiter,
			}
			got, got1 := l.calLeftCapacity(tt.args.b, tt.args.amount, tt.args.capacity, tt.args.period, tt.args.rate, tt.args.nowTimestamp)
			if got != tt.want {
				t.Errorf("calLeftCapacity() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calLeftCapacity() got1 = %v, want %v", got1, tt.want1)
			}
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
