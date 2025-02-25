package dstruct

import (
	"testing"
)

func TestIntStack(t *testing.T) {
	stack := NewStack[int](true)
	if stack.Size() != 0 {
		t.Errorf("stack size[%d] is not equal 0", stack.Size())
	}

	_, ok11 := stack.Peek()
	if ok11 {
		t.Errorf("stack peek is ok[%t]", ok11)
	}

	_, ok22 := stack.Pop()
	if ok22 {
		t.Errorf("stack pop is ok[%t]", ok22)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if stack.Size() != 4 {
		t.Errorf("stack size[%d] is not equal 4", stack.Size())
	}

	peekVal, ok1 := stack.Peek()
	if !ok1 || peekVal != 4 {
		t.Errorf("stack peek[%d] is not equal 4, or peek is not ok[%t]", peekVal, ok1)
	}

	popVal, ok2 := stack.Pop()
	if !ok2 || popVal != 4 {
		t.Errorf("stack pop[%d] is not equal 4, or pop is not ok[%t]", peekVal, ok2)
	}

	stack.Push(5)
	stack.Push(6)

	peekVal, ok1 = stack.Peek()
	if !ok1 || peekVal != 6 {
		t.Errorf("stack peek[%d] is not equal 6, or peek is not ok[%t]", peekVal, ok1)
	}

	for !stack.IsEmpty() {
		stack.Pop()
	}

	if stack.Size() != 0 {
		t.Errorf("stack size[%d] is not equal 0", stack.Size())
	}
}

func TestStringStack(t *testing.T) {
	stack := NewStack[string](false)
	if stack.Size() != 0 {
		t.Errorf("stack size[%d] is not equal 0", stack.Size())
	}

	_, ok11 := stack.Peek()
	if ok11 {
		t.Errorf("stack peek is ok[%t]", ok11)
	}

	_, ok22 := stack.Pop()
	if ok22 {
		t.Errorf("stack pop is ok[%t]", ok22)
	}

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	if stack.Size() != 3 {
		t.Errorf("stack size[%d] is not equal 3", stack.Size())
	}

	peekVal, ok1 := stack.Peek()
	if !ok1 || peekVal != "c" {
		t.Errorf("stack peek[%s] is not equal c, or peek is not ok[%t]", peekVal, ok1)
	}

	popVal, ok2 := stack.Pop()
	if !ok2 || popVal != "c" {
		t.Errorf("stack pop[%s] is not equal c, or peek is not ok[%t]", peekVal, ok2)
	}

	stack.Push("d")
	stack.Push("e")

	peekVal, ok1 = stack.Peek()
	if !ok1 || peekVal != "e" {
		t.Errorf("stack peek[%s] is not equal e, or peek is not ok[%t]", peekVal, ok1)
	}

	for !stack.IsEmpty() {
		stack.Pop()
	}

	if stack.Size() != 0 {
		t.Errorf("stack size[%d] is not equal 0", stack.Size())
	}

}
