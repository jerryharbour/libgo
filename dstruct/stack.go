package dstruct

import (
	"sync"
)

type Stack[T DStructElemType] struct {
	data []T
	len  int
	lock *sync.Mutex
}

func NewStack[T DStructElemType](threadSafe bool) *Stack[T] {
	s := &Stack[T]{
		data: make([]T, 0),
		len:  0,
		lock: nil,
	}

	if threadSafe {
		s.lock = &sync.Mutex{}
	}

	return s
}

func (s *Stack[T]) Push(elem T) {
	if s.lock != nil {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	if s.len >= len(s.data) {
		newCapacity := len(s.data) * 2
		if newCapacity == 0 {
			newCapacity = 1
		}

		dataTmp := make([]T, newCapacity)
		copy(dataTmp, s.data[:s.len])
		s.data = dataTmp
	}

	s.data[s.len] = elem
	s.len++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.lock != nil {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	if s.len == 0 {
		var zero T
		return zero, false
	}

	value := s.data[s.len-1]
	s.len--
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.lock != nil {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	if s.len == 0 {
		var zero T
		return zero, false
	}

	return s.data[s.len-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	if s.lock != nil {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	return s.len == 0
}

func (s *Stack[T]) Size() int {
	if s.lock != nil {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	return s.len
}
