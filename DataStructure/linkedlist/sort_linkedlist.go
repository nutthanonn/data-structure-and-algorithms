package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Val  int
	next *Node
}

type linkedlist struct {
	head *Node
	size int
}

func (l *linkedlist) Insert(val int) {
	new_node := Node{}
	new_node.Val = val
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

func (l *linkedlist) SortLinkedlist() {
	ptr := l.head
	data_arr := []int{}
	for i := 0; i < l.size; i++ {
		data_arr = append(data_arr, ptr.Val)
		ptr = ptr.next
	}
	sort.Slice(data_arr, func(i, j int) bool { return data_arr[i] < data_arr[j] })
	l.size = 0
	for i := 0; i < len(data_arr); i++ {
		l.Insert(data_arr[i])
	}
}

func (l *linkedlist) getLinkedlist() {
	ptr := l.head
	arr := []int{}
	for i := 0; i < l.size; i++ {
		arr = append(arr, ptr.Val)
		ptr = ptr.next
	}
	fmt.Println(arr)
}

func main() {
	l := linkedlist{}
	l.Insert(4)
	l.Insert(34)
	l.Insert(21)
	l.Insert(7)

	fmt.Println("Before Sort Linkedlist: ")
	l.getLinkedlist()

	l.SortLinkedlist()

	fmt.Println("After Sort Linkedlist: ")
	l.getLinkedlist()

}
