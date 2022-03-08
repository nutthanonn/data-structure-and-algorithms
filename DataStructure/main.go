package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type linkedlist struct {
	head *Node
	size int
}

func (l *linkedlist) Insert(val int) {
	new_node := Node{}
	new_node.data = val

	if l.size == 0 {
		l.head = &new_node
		l.size++
		return
	}

	ptr := l.head
	for i := 0; i < l.size; i++ {
		if ptr.next == nil {
			ptr.next = &new_node
			l.size++
			return
		}
		ptr = ptr.next
	}
}

func (l *linkedlist) getAllNode() {
	ptr := l.head
	for i := 0; i < l.size; i++ {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}

func (l *linkedlist) getPosNode(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}

	for i := 0; i < l.size; i++ {
		if i == pos {
			return ptr
		}
		ptr = ptr.next
	}

	return ptr
}

func (l *linkedlist) swapNode(pos1, pos2 int) {
	prevNode := l.getPosNode(pos1 - 1)
	node_1 := l.getPosNode(pos1)
	node_2 := l.getPosNode(pos2)
	temp := node_2.next
	node_2.next = prevNode.next
	prevNode.next = node_1.next
	node_1.next = temp

}

func main() {
	l := linkedlist{}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	l.getAllNode()
	fmt.Println("-------")
	l.swapNode(1, 2)
	l.getAllNode()

}
