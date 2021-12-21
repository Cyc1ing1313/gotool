package queue

import "fmt"

type Queue[T any] struct {
	queue []T
}

func New[T any]() *Queue[T]{
	return &Queue[T]{
		queue: []T{},
	}
}

func (q *Queue[T]) Push(x T) {
	q.queue = append(q.queue, x)
}

func (q *Queue[T]) Pop() (T, error) {
	if len(q.queue) == 0 {
		return *new(T), fmt.Errorf("pop empty queue")
	}
	x := q.queue[0]
	q.queue = q.queue[1:]
	return x, nil
}

func (q *Queue[T]) Front() (T,error){
	if len(q.queue) == 0{
		return *new(T),fmt.Errorf("front empty queue")
	}
	return q.queue[0],nil
}
