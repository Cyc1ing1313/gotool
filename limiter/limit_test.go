package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	var limit = NewLimiter(10)
	for range [5]struct{}{} {
		for{
			if limit.aquire() {
				fmt.Printf("do work\n")
			} else {
				fmt.Printf("drop\n")
			}
			time.Sleep(time.Millisecond * 100)
		}	
	}

	select {}

}
