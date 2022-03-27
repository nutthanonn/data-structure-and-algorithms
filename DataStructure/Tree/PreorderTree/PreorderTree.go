package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	new_node := Node{}
	new_node.Val = val

	if n.Val > val {
		if n.Left == nil {
			n.Left = &new_node
		} else {
			n.Left.Insert(val)
		}
	} else if n.Val < val {
		if n.Right == nil {
			n.Right = &new_node
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}

	if n.Val > val {
		return n.Left.Search(val)
	} else if n.Val < val {
		return n.Right.Search(val)
	}

	return true
}

func main() {
	t := Node{Val: 200}
	t.Insert(10)
	t.Insert(2304)
	t.Insert(123)
	t.Insert(2345)
	t.Insert(31)

	fmt.Println(t.Search(31))
	fmt.Println(t.Search(10))
	fmt.Println(t.Search(2304))
	fmt.Println(t.Search(230))
}
