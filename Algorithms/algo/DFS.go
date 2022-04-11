package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (this *Node) Insert(val int) {
	new_node := &Node{}
	new_node.Val = val

	if val > this.Val {
		if this.Right == nil {
			this.Right = new_node
		} else {
			this.Right.Insert(val)
		}
	} else if val < this.Val {
		if this.Left == nil {
			this.Left = new_node
		} else {
			this.Left.Insert(val)
		}
	}
}

func (this *Node) DFS() {
	if this == nil {
		return
	}

	this.Left.DFS()
	fmt.Println(this.Val)
	this.Right.DFS()
}

// ########################################################################################################################################################################

func main() {
	t := &Node{Val: 100}
	t.Insert(110)
	t.Insert(110)
	t.Insert(120)
	t.Insert(105)

	t.Insert(50)
	t.Insert(20)
	t.Insert(90)

	// t.DFS()
}
