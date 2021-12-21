package stack

import "fmt"

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T]{
	return &Stack[T]{
		data: []T{},
	}
}

func (s *Stack[T]) push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) pop() (T,error) {
	if len(s.data) == 0 {
		return *new(T), fmt.Errorf("pop empty stack")
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[0:len(s.data)-1]
	return x,nil
}

func (s *Stack[T]) peek() (T,error) {
	if len(s.data) == 0 {
		return *new(T), fmt.Errorf("peek empty stack")
	}
	x := s.data[len(s.data)-1]
	return x,nil
}