package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (this *Node) Insert(val int) {
	new_node := &Node{Val: val}

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

func (n *Node) BFS() {
	if n == nil {
		return
	}
	q := []*Node{}
	q = append(q, n)

	for len(q) != 0 {
		ptr := q[0]
		fmt.Println(ptr.Val)
		q = q[1:]

		if ptr.Left != nil {
			q = append(q, ptr.Left)
		}

		if ptr.Right != nil {
			q = append(q, ptr.Right)
		}
	}
}

func main() {
	t := &Node{Val: 100}
	t.Insert(110)
	t.Insert(120)
	t.Insert(105)

	t.Insert(50)
	t.Insert(20)
	t.Insert(90)
	t.BFS()
}
