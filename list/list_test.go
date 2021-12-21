package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := New[int](true)
	list.Add(1)
	list.Add(2)
	list.Map(func(x int) int {
		return 3 * x
	}).Filter(func(x int) bool {
		return x > 3
	}).Foreach(func(i, t int) {
		fmt.Println(i, t)
	})

}
