package main

type Queue[T any] struct {
	elements []T
}

// Generare a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue a new element into the queue
func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

// Next removes and returns the first element from the queue
func (q *Queue[T]) Next() (T, bool) {
	var zero T
	if len(q.elements) == 0 {
		return zero, false
	}
	next := q.elements[0]
	q.elements = q.elements[1:]
	return next, true
}

// Len return the current lenght of the queue
func (q *Queue[T]) Len() int {
	return len(q.elements)
}

// IsEmpty return true if the queue has no elements
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
