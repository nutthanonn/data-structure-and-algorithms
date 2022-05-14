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
func (this *Node) LevelOrder() {
	if this == nil {
		return
	}

	q := []*Node{}
	q = append(q, this) //push

	for len(q) != 0 {
		ptr := q[0]
		fmt.Println(ptr.Val)
		q = q[1:] //pop

		if ptr.Left != nil {
			q = append(q, ptr.Left)
		}

		if ptr.Right != nil {
			q = append(q, ptr.Right)
		}
	}
}

func (this *Node) printLevelOrderLinebyLine() {
	if this == nil {
		return
	}

	q := []*Node{}
	q = append(q, this)
	for len(q) != 0 {
		count := len(q)
		for count > 0 {
			temp := q[0]
			q = q[1:]
			fmt.Printf("%d  ", temp.Val)
			if temp.Left != nil {
				q = append(q, temp.Left)
			}

			if temp.Right != nil {
				q = append(q, temp.Right)
			}
			count--
		}
		fmt.Printf("\n")
	}

}

func main() {
	t := &Node{Val: 100}
	t.Insert(150)
	t.Insert(50)

	t.LevelOrder()
}
