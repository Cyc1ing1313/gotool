package hashmap

import (
	"sync"
)

type HashMap[T comparable, P any] struct {
	data map[T]P
	lock sync.RWMutex
	safe bool
	sz int
}



func New[T comparable, P any](threadSafe bool) *HashMap[T, P] {
	t := &HashMap[T, P]{
		data: make(map[T]P),
		safe: threadSafe,
	}
	if threadSafe {
		t.lock = sync.RWMutex{}
	}
	return t
}

func (m *HashMap[T,P]) Foreach(f func(T,P)) {
	if m.safe{
		m.lock.RLock()
		for k,v:=range m.data{
			f(k,v)
		}
		m.lock.RUnlock()
	}else{
		for k,v:=range m.data{
			f(k,v)
		}
	}
}

func (m *HashMap[T, P]) Map(key T, f func(P) P) *HashMap[T, P] {
	if m.safe {
		tmap := map[T]P{}
		m.lock.RLock()
		for k, v := range m.data {
			tmap[k] = f(v)
		}
		m.lock.RUnlock()
		m.lock.Lock()
		m.data = tmap
		m.lock.Unlock()
		return m
	}
	for k, v := range m.data {
		m.data[k] = f(v)
	}
	return m
}

func (m *HashMap[T, P]) Filter(f func(T, P) bool) *HashMap[T, P] {
	t := map[T]P{}
	if m.safe {
		m.lock.RLock()
		for k, v := range m.data {
			if f(k, v) {
				t[k] = v
			}
		}
		m.lock.RUnlock()
		m.lock.Lock()
		m.data = t
		m.lock.Unlock()
		return m
	}
	for k, v := range m.data {
		if f(k, v) {
			t[k] = v
		}
	}
	m.data = t
	return m
}

func (m *HashMap[T, P]) Get(k T) (P, bool) {
	if m.safe {
		m.lock.RLock()
		val, ok := m.data[k]
		m.lock.RUnlock()
		return val, ok
	}
	val, ok := m.data[k]
	return val, ok
}

func (m *HashMap[T, P]) Set(k T, v P) {
	if m.safe {
		m.lock.Lock()
		m.data[k] = v
		m.lock.Unlock()
		return
	}
	m.data[k] = v
}

func (m *HashMap[T, P]) Del(k T) {
	if m.safe {
		m.lock.Lock()
		delete(m.data, k)
		m.lock.Unlock()
		return
	}
	delete(m.data, k)
}

func (m *HashMap[T, P]) Size() int {
	if m.safe {
		m.lock.RLock()
		defer m.lock.RUnlock()
		return len(m.data)
	}
	return len(m.data)
}
