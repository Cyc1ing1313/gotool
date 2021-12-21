package limiter

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Limiter struct {
	qps int64
}

func NewLimiter(qps int64) *Limiter {
	l := &Limiter{
		qps: qps,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("%+v", err)
			}
		}()
		for range time.Tick(time.Second) {
			atomic.StoreInt64(&(l.qps), qps)
		}
	}()
	return l
}

func (l *Limiter) aquire() bool {
	if atomic.AddInt64(&l.qps, -1) < 0 {
		atomic.AddInt64(&l.qps, 1)
		return false
	} else {
		return true
	}
}
