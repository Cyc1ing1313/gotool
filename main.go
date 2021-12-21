package main

import "fmt"

func hello[T any](x T) {
	fmt.Println(x)
}

func main() {
	a := make([]int, 10)
	for i:=9;i<100;i++{
		a[i] = i
	}
}
