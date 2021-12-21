package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	m := New[int, int](true)
	m.Set(1, 1)
	m.Set(2, 3)
	m.Map(1, func(x int) int {
		return 2 * x
	}).Filter(func(k, v int) bool {
		return v > 2
	}).Foreach(func(k, v int) {
		fmt.Println(k, v)
	})
}
