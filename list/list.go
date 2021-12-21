package list

import "sync"

type List[T any] struct {
	data []T
	safe bool
	lock sync.RWMutex
}

func New[T any](threadSafe bool) *List[T] {
	t := &List[T]{
		data: make([]T, 0),
		safe: threadSafe,
	}
	if threadSafe {
		t.lock = sync.RWMutex{}
	}
	return t
}

func (l *List[T]) Add(x T) {
	if l.safe {
		l.lock.Lock()
		l.data = append(l.data, x)
		l.lock.Unlock()
	} else {
		l.data = append(l.data, x)
	}
}

func (l *List[T]) Del(index int) {
	if l.safe {
		l.lock.Lock()
		if index < 0 || index >= len(l.data) {
			return
		}
		l.data = append(l.data[:index], l.data[index+1:]...)
		l.lock.Unlock()
	} else {
		if index < 0 || index >= len(l.data) {
			return
		}
		l.data = append(l.data[:index], l.data[index+1:]...)
	}
}

func (l *List[T]) Get(index int) (T, bool) {
	if l.safe {
		l.lock.RLock()
		defer l.lock.RUnlock()
		if index < 0 || index >= len(l.data) {
			return *new(T), false
		}
		return l.data[index], true
	} else {
		if index < 0 || index >= len(l.data) {
			return *new(T), false
		}
		return l.data[index], true
	}
}

func (l *List[T]) Map(f func(T) T) *List[T] {
	if l.safe {
		l.lock.Lock()
		for i := 0; i < len(l.data); i++ {
			l.data[i] = f(l.data[i])
		}
		l.lock.Unlock()
		return l
	}
	for i := 0; i < len(l.data); i++ {
		l.data[i] = f(l.data[i])
	}
	return l
}

func (l *List[T]) Filter(f func(T) bool) *List[T] {
	if l.safe {
		l.lock.RLock()
		t := make([]T, 0)
		for i := 0; i < len(l.data); i++ {
			if f(l.data[i]) {
				t = append(t, l.data[i])
			}
		}
		l.lock.RUnlock()
		l.lock.Lock()
		l.data = t
		l.lock.Unlock()
		return l
	} else {
		t := make([]T, 0)
		for i := 0; i < len(l.data); i++ {
			if f(l.data[i]) {
				t = append(t, l.data[i])
			}
		}
		l.data = t
		return l
	}
}

func (l *List[T]) Foreach(f func(int, T)) {
	if l.safe {
		l.lock.RLock()
		for i, v := range l.data {
			f(i, v)
		}
		l.lock.RUnlock()
	} else {
		for i, v := range l.data {
			f(i, v)
		}
	}
}
