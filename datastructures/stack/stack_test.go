package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	if s.head != nil || s.length != 0 {
		t.Fatal("Head of the stack must be nil and length == 0 when just started instance")
	}
	s.Push(1)
	v := s.Peek()
	if v.Get() != 1 || s.Size() != 1 {
		t.Fatal("Peek must return the last element and Size must be 1")
	}
	s.Push(2)
	v, ok := s.Pop()
	if !ok {
		t.Fatal("Pop must be able to remove the last element, not ok returned")
	}
	if v.Get() != 2 {
		t.Fatal("Pop must be able to remove the last element, value different of 2 returned")
	}
	if s.Size() != 1 {
		t.Fatal("Pop must be able to remove the last element, length different of 1 returned")
	}
}
