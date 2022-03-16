package main

import "fmt"

type Node struct {
	val   int
	Left  *Node
	Right *Node
}

type InorderArr struct {
	arr  []int
	size int
}

func (n *Node) Insert(val int) {
	new_node := Node{}
	new_node.val = val
	if n.val > val {
		if n.Left == nil {
			n.Left = &new_node
		} else {
			n.Left.Insert(val)
		}
	} else if n.val < val {
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

	if n.val > val {
		return n.Left.Search(val)
	} else if n.val < val {
		return n.Right.Search(val)
	}

	return true
}

//Method Inorder
func (i *InorderArr) InorderTree(n *Node) {
	if n == nil {
		return
	}

	i.InorderTree(n.Left)
	i.arr = append(i.arr, n.val)
	i.size++
	i.InorderTree(n.Right)
}

func main() {
	tree := &Node{val: 25}
	inorder := &InorderArr{}
	node := []int{15, 50, 10, 22, 35, 70, 4, 12, 18, 24, 31, 44, 66, 90}
	for i := 0; i < len(node); i++ {
		tree.Insert(node[i])
	}
	inorder.InorderTree(tree)
	fmt.Println(inorder)

}
