package dstruct

import (
	"sync"
)

type FifoRingQ[T DStructElemType] struct {
	data   []T
	len    int
	head   int
	tail   int
	maxLen int
	lock   *sync.Mutex
}

func newFifoRingQ[T DStructElemType](maxLen int, threadSafe bool) *FifoRingQ[T] {
	if maxLen <= 0 {
		maxLen = 10 // default maxLen
	}

	q := &FifoRingQ[T]{
		data:   make([]T, maxLen),
		len:    0,
		head:   0,
		tail:   0,
		maxLen: maxLen,
		lock:   nil,
	}

	if threadSafe {
		q.lock = &sync.Mutex{}
	}

	return q
}

func NewFifoRingQ[T DStructElemType](maxLen int) *FifoRingQ[T] {
	return newFifoRingQ[T](maxLen, false)
}

func NewSafeFifoRingQ[T DStructElemType](maxLen int) *FifoRingQ[T] {
	return newFifoRingQ[T](maxLen, true)
}

func (f *FifoRingQ[T]) Push(elem T) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	if f.maxLen == 0 {
		// if max len is 0, it is not ring queue
		return
	}

	if f.len == f.maxLen {
		// ring is full, cover old element
		f.head = (f.head + 1) % f.maxLen
		f.len--
	}

	f.data[f.tail] = elem
	f.tail = (f.tail + 1) % f.maxLen
	f.len++
}

func (f *FifoRingQ[T]) Pop() (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	if f.len == 0 {
		var zero T
		return zero, false
	}

	elem := f.data[f.head]
	// remove element
	f.head = (f.head + 1) % f.maxLen
	f.len--
	return elem, true
}

func (f *FifoRingQ[T]) PeekFront() (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	if f.len == 0 {
		var zero T
		return zero, false
	}

	return f.data[f.head], true
}

func (f *FifoRingQ[T]) PeekBack() (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	if f.len == 0 {
		var zero T
		return zero, false
	}

	return f.data[(f.tail-1+f.maxLen)%f.maxLen], true
}

func (f *FifoRingQ[T]) PeekIdx(idx int) (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	if f.len == 0 || idx < 0 || idx >= f.len {
		var zero T
		return zero, false
	}

	return f.data[(f.head+idx)%f.maxLen], true
}

func (f *FifoRingQ[T]) IsEmpty() bool {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	return f.len == 0
}

func (f *FifoRingQ[T]) Size() int {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	return f.len
}
