/*

Floyd’s Cycle Detection Algorithm
- detection Cycle linkedlist


*/

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type linkedlist struct {
	head *Node
	size int
}

func (this *linkedlist) Insert(val int) {
	new_node := &Node{Val: val}
	if this.head == nil {
		this.head = new_node
		this.size++
		return
	}

	ptr := this.head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = new_node
	this.size++
}
func (this *linkedlist) getData() {
	ptr := this.head
	for ptr != nil {
		fmt.Println(ptr.Val)
		ptr = ptr.Next
	}
}
func (this *linkedlist) makeCycleLinkedlist(index int) {
	if index > this.size {
		return
	}

	ptr := this.head
	end_node := this.head

	for end_node.Next != nil {
		end_node = end_node.Next
	}

	for i := 0; i < index; i++ {
		ptr = ptr.Next
	}

	end_node.Next = ptr
}

func (this *linkedlist) cycleDetection() bool {
	q, p := this.head, this.head

	for q != nil && p.Next != nil {
		q = q.Next
		p = p.Next.Next

		if q == p {
			return true
		}
	}

	return false
}

func main() {
	l := &linkedlist{}
	l.Insert(100)
	l.Insert(20)
	l.Insert(30)
	l.Insert(40)
	// l.getData()
	l.makeCycleLinkedlist(1)
	fmt.Println(l.cycleDetection())
}
