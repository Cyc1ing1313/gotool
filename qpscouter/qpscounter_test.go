package qpscouter

import (
	"testing"
	"time"
)

func TestQpsCounter(t *testing.T) {
	qpsCounter := New()
	for i := 0; i < 1000000; i++ {
		qpsCounter.count()
		time.Sleep(time.Second)
		t.Log(qpsCounter.report())
	}
}
