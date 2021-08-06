package memo

import (
	"error"
	"sync"
)

type Memo struct {
	f Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

var mu sync.Mutex

func (m *Memo) Get(key string) (interface{}, error) {
	mu.Lock()
	defer mu.Unlock()

	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}