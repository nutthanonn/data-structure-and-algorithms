package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var count int

func (n *Node) Insert(val int) {
	node := Node{}
	node.Val = val
	if n.Val < val {
		//right child
		if n.Right == nil {
			n.Right = &node
		} else {
			n.Right.Insert(val)
		}
	} else if n.Val > val {
		//left
		if n.Left == nil {
			n.Left = &node
		} else {
			n.Left.Insert(val)
		}
	}
}

func (n *Node) Search(val int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Val < val {
		return n.Right.Search(val)
	} else if n.Val > val {
		return n.Left.Search(val)
	}

	return true
}

func main() {
	tree := &Node{Val: 100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(276))
	fmt.Println(count)
}
