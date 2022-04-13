package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Array struct {
	s []int
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

func (this *Array) ProOrder(n *Node) {
	if n == nil {
		return
	}

	this.s = append(this.s, n.Val)
	this.ProOrder(n.Left)
	this.ProOrder(n.Right)
}

func main() {
	t := &Node{Val: 100}
	t.Insert(50)
	t.Insert(90)
	t.Insert(10)

	t.Insert(150)
	t.Insert(190)
	t.Insert(110)

	s := &Array{}
	s.ProOrder(t)
	fmt.Println(s.s)

}
