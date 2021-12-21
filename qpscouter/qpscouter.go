package qpscouter

import (
	"fmt"
	"sync/atomic"
	"time"
)

type QpsCouter struct {
	qpsMax   int64
	qpsMin   int64
	qpsAvg   int64
	lifeTime int64
	qps      int64
}

func New() *QpsCouter {
	q := &QpsCouter{
		qpsMax:   0,
		qpsMin:   0,
		qpsAvg:   0,
		lifeTime: 0,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("qpscouter panic: %v", err)
			}
		}()
		for range time.Tick(time.Second) {
			qps := atomic.SwapInt64(&q.qps, 0)
			if q.qpsMax != 0 {
				q.qpsMax = int64Max(q.qpsMax, qps)
			} else {
				q.qpsMax = qps
			}

			if q.qpsMin != 0 {
				q.qpsMin = int64Min(q.qpsMin, qps)
			} else {
				q.qpsMin = qps
			}
			if q.lifeTime != 0 {
				q.qpsAvg = int64((float64(q.qpsAvg) + int64Sub(qps, q.lifeTime)) / (1 + int64Sub(1, q.lifeTime)))
			} else {
				q.qpsAvg = qps
			}
			atomic.AddInt64(&q.lifeTime, 1)
		}
	}()
	return q
}

func (c *QpsCouter) count() {
	atomic.AddInt64(&c.qps, 1)
}

func (c *QpsCouter) report() string {
	return fmt.Sprintf("qps:%d,max:%d,min:%d,avg:%d", c.qps, c.qpsMax, c.qpsMin, c.qpsAvg)
}

func (c *QpsCouter) getQpsMax() int64 {
	return c.qpsMax
}

func (c *QpsCouter) getQpsMin() int64 {
	return c.qpsMin
}

func (c *QpsCouter) getQpsAvg() int64 {
	return c.qpsAvg
}

func int64Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func int64Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func int64Sub(a,b int64) float64{
	return float64(a)/float64(b)
}