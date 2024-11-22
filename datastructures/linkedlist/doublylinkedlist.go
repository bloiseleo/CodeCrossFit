package linkedlist

import "fmt"

type LinkedListNode struct {
	next     *LinkedListNode
	previous *LinkedListNode
	Value    interface{}
}

type DoublyLinkedList struct {
	head   *LinkedListNode
	tail   *LinkedListNode
	length int
}

func LinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

/*
Transform the LinkedList into dynamic slice. O(n)
*/
func (self *DoublyLinkedList) ToArray() []interface{} {
	arr := make([]interface{}, 0)
	if self.length == 0 {
		return arr
	}
	node := self.head
	for i := 0; i < self.length; i++ {
		if node == nil {
			break
		}
		arr = append(arr, node.Value)
		node = node.next
	}
	return arr
}

/*
Inserts a node at the beginning of the array. O(1)
*/
func (self *DoublyLinkedList) Unshfit(node *LinkedListNode) {
	node.previous = nil
	if self.length == 0 {
		node.next = nil
		node.previous = nil
		self.head = node
		self.tail = node
		self.length++
		return
	}
	oldHead := self.head
	oldHead.previous = node
	node.next = oldHead
	self.head = node
	self.length++
}

/*
Inserts a node at the end of the array. O(1)
*/
func (self *DoublyLinkedList) Append(node *LinkedListNode) {
	if self.length == 0 {
		node.next = nil
		node.previous = nil
		self.head = node
		self.tail = node
		self.length++
		return
	}
	oldTail := self.tail
	node.previous = oldTail
	node.next = nil
	oldTail.next = node
	self.tail = node
	self.length++
}

/*
Insert a node inside the linked list at [position] defined. O(n)
*/
func (self *DoublyLinkedList) Insert(node *LinkedListNode, position int) {
	if position == 0 {
		self.Unshfit(node)
		return
	}
	if position == self.length-1 {
		self.Append(node)
		return
	}
	if position >= self.length || position < 0 {
		panic(fmt.Sprintf("Index [%d] out of bounds", position))
	}
	existingNode := self.head
	for i := 0; i < position; i++ {
		existingNode = existingNode.next
	}
	previous := existingNode.previous
	previous.next = node
	node.previous = previous
	node.next = existingNode
	existingNode.previous = node
	self.length++
}

func (self *DoublyLinkedList) Delete(position int) {
	if position < 0 || position >= self.length {
		panic(fmt.Sprintf("Index [%d] out of bounds", position))
	}
	existingNode := self.head
	for i := 0; i < position; i++ {
		existingNode = existingNode.next
	}
	previous := existingNode.previous
	next := existingNode.next
	existingNode.previous = nil
	existingNode.next = nil
	if next != nil {
		next.previous = previous
	}
	if previous != nil {
		previous.next = next
	}
	self.length--
}
