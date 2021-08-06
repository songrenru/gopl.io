package memo

import (
	"error"
	"sync"
)

type entry struct {
	res result
	ready chan struct{}
}

type Memo struct {
	f Func
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

var mu sync.Mutex

func (m *Memo) Get(key string) (interface{}, error) {
	mu.Lock()
	e := m.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		mu.Unlock()

		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	} else {
		mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}