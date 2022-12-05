package util

import (
	"errors"
	"strconv"
)

type Stack[T any] []T

func (s *Stack[T]) PushAll(elements Stack[T]) {
	*s = append(elements, *s...)
}

func (s *Stack[T]) Push(element T) {
	*s = append([]T{element}, *s...)
}

func (s *Stack[T]) PushLast(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Pop() (T, error) {
	element, err := s.Peek()
	if err != nil {
		return element, err
	}
	*s = (*s)[1:]
	return element, nil
}

func (s *Stack[T]) PopMultiple(amount int) (Stack[T], error) {
	if s.Len() < amount {
		return nil, errors.New("stack does not contain " + strconv.Itoa(amount) + " elements")
	}
	result := append(Stack[T]{}, (*s)[0:amount]...)
	*s = (*s)[amount:]

	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Len() == 0 {
		var empty T
		return empty, errors.New("stack is empty")
	}
	return (*s)[0], nil
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
