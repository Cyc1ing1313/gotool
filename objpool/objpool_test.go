package objpool

import (
	"testing"
)

type Man struct {
	Name string
	Age  int
}

func TestObjPool(t *testing.T) {
	pool := New[Man](1,func(t *Man) {
		t.Age = 0
		t.Name = "hhh"
	})
	man := pool.Get()
	man.Age = 10
	t.Logf("%#v", man)
	pool.Put(man)
	man = pool.Get()
	t.Logf("%#v", man)
}
