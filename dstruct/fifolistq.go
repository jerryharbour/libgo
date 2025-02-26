package dstruct

import (
	"sync"
	"container/list"
)

type FifoListQ[T DStructElemType] struct {
	data *list.List
	maxLen int
	lock *sync.Mutex
}

func newFifoListQ[T DStructElemType](maxLen int, threadSafe bool) *FifoListQ[T] {
	q := &FifoListQ[T]{
		data: list.New(),
		maxLen: maxLen,
		lock: nil,
	}

	if threadSafe {
		q.lock = &sync.Mutex{}
	}

	return q
}

func NewFifoListQ[T DStructElemType]() *FifoListQ[T] {
	return newFifoListQ[T](0, false)
}

func NewSafeFifoListQ[T DStructElemType]() *FifoListQ[T] {
	return newFifoListQ[T](0, true)
}

func NewLimitFifoListQ[T DStructElemType](maxLen int) *FifoListQ[T] {
	return newFifoListQ[T](maxLen, false)
}

func NewSafeLimitFifoListQ[T DStructElemType](maxLen int) *FifoListQ[T] {
	return newFifoListQ[T](maxLen, true)
}

func (f *FifoListQ[T]) Push(elem T) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	if f.maxLen > 0 && f.data.Len() >= f.maxLen {
		f.pop()
	}

	f.data.PushBack(elem)
}

func (f *FifoListQ[T]) Pop() (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	return f.pop()
}

func (f *FifoListQ[T]) pop() (T, bool) {
	var zero T

	if  f.data.Len() == 0 {
		return zero, false
	}

	elem := f.data.Front()
	if elem == nil {
		return zero, false
	}

	f.data.Remove(elem)
	return elem.Value.(T), true
}

func (f *FifoListQ[T]) PeekFront() (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	var zero T
	if  f.data.Len() == 0 {
		return zero, false
	}

	elem := f.data.Front()
	if elem == nil {
		return zero, false
	}

	return elem.Value.(T), true
}

func (f *FifoListQ[T]) PeekBack() (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	var zero T
	if  f.data.Len() == 0 {
		return zero, false
	}

	elem := f.data.Back()
	if elem == nil {
		return zero, false
	}

	return elem.Value.(T), true
}

func (f *FifoListQ[T]) PeekIdx(idx int) (T, bool) {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	var zero T
	if idx < 0 || idx >= f.data.Len() {
		return zero, false
	}

	var elem *list.Element
	elem = nil
	bNext := true
	stepLen := idx + 1
	if idx < f.data.Len()/2 {
		elem = f.data.Front()
	} else {
		elem = f.data.Back()
		bNext = false
		stepLen = f.data.Len() - idx
	}

	if elem == nil {
		return zero, false
	}

	for i := 0; i < stepLen; i++ {
		if bNext {
			elem = elem.Next()
		} else {
			elem = elem.Prev()
		}

		if elem == nil {
			break
		}
	}

	if elem == nil {
		return zero, false
	}

	return elem.Value.(T), true
}

func (f *FifoListQ[T]) IsEmpty() bool {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	return f.data.Len() == 0
}

func (f *FifoListQ[T]) Size() int {
	if f.lock != nil {
		f.lock.Lock()
		defer f.lock.Unlock()
	}

	return f.data.Len()
}