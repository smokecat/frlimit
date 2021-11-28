package frlimit

type Adapter interface {
	Get(key string) (int, error)
	Set(key string, value int) error
	Incr(key string) (int, error)
	Del(key string) error
}
