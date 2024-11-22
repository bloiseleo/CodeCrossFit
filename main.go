package main

import (
	"fmt"

	"github.com/bloiseleo/CodeCrossFit/datastructures/linkedlist"
)

func main() {
	linkedList := linkedlist.LinkedList()
	for i := 0; i < 2; i++ {
		linkedList.Append(&linkedlist.LinkedListNode{
			Value: i,
		})
	}
	linkedList.Delete(1)
	linkedList.Delete(0)
	arr := linkedList.ToArray()
	fmt.Println(arr...)
}
