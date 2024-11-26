package stack

type stackNode struct {
	value interface{}
	prev  *stackNode
}

type Stack struct {
	head   *stackNode
	length int
}

func (self *stackNode) Get() interface{} {
	return self.value
}

/*
Retrieve the size of the stack. O(1)
*/
func (self *Stack) Size() int {
	return self.length
}

/*
Peek the first element of the Stack. O(1)
*/
func (self *Stack) Peek() *stackNode {
	return self.head
}

/*
Push an element on top of the stack. O(1)
*/
func (self *Stack) Push(value interface{}) {
	if self.head == nil {
		self.head = &stackNode{
			value: value,
			prev:  nil,
		}
		self.length++
		return
	}
	self.head = &stackNode{
		value: value,
		prev:  self.head,
	}
	self.length++
}

/*
Remove the last element of the stack. O(1)
*/
func (self *Stack) Pop() (*stackNode, bool) {
	if self.length == 0 {
		return nil, false
	}
	node := self.head
	self.head = self.head.prev
	self.length--
	return node, true
}
