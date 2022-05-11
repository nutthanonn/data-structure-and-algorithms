package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (this *Node) Insert(val int) {
	node := Node{}
	node.Val = val
	if this.Val < val {
		if this.Right == nil {
			this.Right = &node
		} else {
			this.Right.Insert(val)
		}
	} else if this.Val > val {
		if this.Left == nil {
			this.Left = &node
		} else {
			this.Left.Insert(val)
		}
	}
}

func (this *Node) FindLeaf() {
	if this.Left == nil && this.Right == nil {
		fmt.Println(this.Val)
		return
	}

	if this.Left != nil {
		this.Left.FindLeaf()
	}

	if this.Right != nil {
		this.Right.FindLeaf()
	}
}

func main() {
	t := &Node{Val: 100}

	t.Insert(50)
	t.Insert(90)
	t.Insert(20)

	t.Insert(120)
	t.Insert(150)
	t.Insert(110)

	t.FindLeaf()
}
