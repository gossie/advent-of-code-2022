package util

import (
	"errors"
)

type queue[T any] []T

func CreateQueue[T any]() queue[T] {
	return make([]T, 0)
}

func (s *queue[T]) Enqueue(element T) {
	*s = append(*s, element)
}

func (s *queue[T]) Dequeue() (T, error) {
	element, err := s.Peek()
	if err != nil {
		return element, err
	}
	*s = (*s)[1:]
	return element, nil
}

func (s *queue[T]) Peek() (T, error) {
	if s.Len() == 0 {
		var empty T
		return empty, errors.New("stack is empty")
	}
	return (*s)[0], nil
}

func (s *queue[T]) Len() int {
	return len(*s)
}

func (s *queue[T]) Empty() bool {
	return s.Len() == 0
}
