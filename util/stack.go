package util

import (
	"errors"
	"strconv"
)

type Stack []string

func (s *Stack) PushAll(elements Stack) {
	*s = append(elements, *s...)
}

func (s *Stack) Push(element string) {
	*s = append([]string{element}, *s...)
}

func (s *Stack) PushLast(element string) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (string, error) {
	element, err := s.Peek()
	if err != nil {
		return element, err
	}
	*s = (*s)[1:]
	return element, nil
}

func (s *Stack) PopMultiple(amount int) (Stack, error) {
	if s.Len() < amount {
		return nil, errors.New("stack does not contain " + strconv.Itoa(amount) + " elements")
	}
	result := append(Stack{}, (*s)[0:amount]...)
	*s = (*s)[amount:]

	return result, nil
}

func (s *Stack) Peek() (string, error) {
	if s.Len() == 0 {
		return "", errors.New("stack is empty")
	}
	return (*s)[0], nil
}

func (s *Stack) Len() int {
	return len(*s)
}
