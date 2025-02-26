package dstruct

import (
	"testing"
)

func TestFifoRingQ(t *testing.T) {
	ringQ := NewFifoRingQ[int](6)
	if ringQ.Size() != 0 {
		t.Errorf("Expected FifoRingQ size to be 0, size %d", ringQ.Size())
	}

	_, peekOk1 := ringQ.PeekFront()
	if peekOk1 {
		t.Errorf("Expected PeekFront to return false, got true")
	}

	_, peekOk1 = ringQ.PeekBack()
	if peekOk1 {
		t.Errorf("Expected PeekBack to return false, got true")
	}

	_, peekOk1 = ringQ.PeekIdx(0)
	if peekOk1 {
		t.Errorf("Expected PeekIdx to return false, got true")
	}

	_, popOk1 := ringQ.Pop()
	if popOk1 {
		t.Errorf("Expected Pop to return false, got true")
	}

	ringQ.Push(1)
	ringQ.Push(2)
	ringQ.Push(3)
	ringQ.Push(4)

	if ringQ.Size() != 4 {
		t.Errorf("Expected FifoRingQ size to be 4, size %d", ringQ.Size())
	}
	// PeekFront
	val, peekOk2 := ringQ.PeekFront()
	if !peekOk2 {
		t.Errorf("Expected PeekFront to return true, got false")
	}
	if val != 1 {
		t.Errorf("Expected PeekFront to return 1, got %d", val)
	}
	// PeekBack
	val, peekOk2 = ringQ.PeekBack()
	if !peekOk2 {
		t.Errorf("Expected PeekBack to return true, got false")
	}
	if val != 4 {
		t.Errorf("Expected PeekBack to return 4, got %d", val)
	}
	// PeekIdx
	val, peekOk2 = ringQ.PeekIdx(2)
	if !peekOk2 {
		t.Errorf("Expected PeekIdx to return true, got false")
	}
	if val != 3 {
		t.Errorf("Expected PeekIdx to return 3, got %d", val)
	}

}
