package objpool

type ObjPool[T any] struct {
	pool  chan T
	maxSz int
	clearFunc func(*T)
}

func New[T any](max int,clear func(*T)) *ObjPool[T] {
	objs := make(chan T,max)
	for i:=0;i<max;i++{
		objs <- *new(T)
	}
	return &ObjPool[T]{
		pool:  objs,
		maxSz: max,
		clearFunc: clear,
	}
}

func (p *ObjPool[T]) Get() *T {
	select {
	case x := <-p.pool:
		return &x
	default:
		return new(T)
	}
}

func (p *ObjPool[T]) Put(x *T) {
	if x == nil {
		return
	}
	p.clearFunc(x)
	select {
	case p.pool <- *x:
		return
	default:
		return
	}
}

func (p *ObjPool[T]) Close() {
	close(p.pool)
}
