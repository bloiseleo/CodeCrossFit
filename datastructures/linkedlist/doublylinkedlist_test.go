package linkedlist

import "testing"

func TestAppendElement(t *testing.T) {
	linkedList := LinkedList()
	linkedList.Append(&LinkedListNode{
		Value: 10,
	})
	linkedList.Append(&LinkedListNode{
		Value: 11,
	})
	linkedList.Append(&LinkedListNode{
		Value: 12,
	})
	if linkedList.tail.Value != 12 {
		t.Fatalf("Last element of the linked list must be 12, %d returned", linkedList.tail.Value)
	}
}

func TestUnshiftElement(t *testing.T) {
	linkedList := LinkedList()
	linkedList.Unshfit(&LinkedListNode{
		Value: 10,
	})
	linkedList.Unshfit(&LinkedListNode{
		Value: 11,
	})
	linkedList.Unshfit(&LinkedListNode{
		Value: 12,
	})
	if linkedList.head.Value != 12 {
		t.Fatalf("First element of the linked list must be 12, %d return", linkedList.head.Value)
	}
}

func TestUnshiftAndAppendElement(t *testing.T) {
	linkedList := LinkedList()
	linkedList.Unshfit(&LinkedListNode{
		Value: 10,
	})
	linkedList.Append(&LinkedListNode{
		Value: 11,
	})
	linkedList.Unshfit(&LinkedListNode{
		Value: 9,
	})
	linkedList.Append(&LinkedListNode{
		Value: 12,
	})
	if linkedList.head.Value != 9 {
		t.Fatalf("First element must be 9, %d returned", linkedList.head.Value)
	}
	if linkedList.tail.Value != 12 {
		t.Fatalf("Last element must be 12, %d returned", linkedList.tail.Value)
	}
}

func TestInsert(t *testing.T) {
	linkedList := LinkedList()
	linkedList.Insert(&LinkedListNode{
		Value: 1,
	}, 0)
	linkedList.Append(&LinkedListNode{
		Value: 2,
	})
	linkedList.Insert(&LinkedListNode{
		Value: 3,
	}, 0)
	if linkedList.length != 3 {
		t.Fatalf("Length of linked list must be 3, %d returned", linkedList.length)
	}
	if linkedList.head.Value != 3 {
		t.Fatalf("Head must be 3, %d returned", linkedList.head.Value)
	}
	if linkedList.tail.Value != 2 {
		t.Fatalf("Tail must be 2, %d returned", linkedList.tail.Value)
	}
}
