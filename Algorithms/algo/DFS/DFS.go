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

//Spanning Tree

type Graph struct {
	graph map[string][]string
}

type Stack struct {
	s []string
}

func newGraph() *Graph {
	return &Graph{graph: make(map[string][]string)}
}

func (this *Graph) Insert(first, last string) {
	this.graph[first] = append(this.graph[first], last)
	this.graph[last] = append(this.graph[last], first)
}

func (this *Stack) DFS_Spanning_Tree(graph map[string][]string, node string) {
	check := true
	for _, v := range this.s {
		if v == node {
			check = false
		}
	}

	if check {
		this.s = append(this.s, node)
		for _, v := range graph[node] {
			this.DFS_Spanning_Tree(graph, v)
		}
	}
}

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

	// ########################################################################################################################################################################

	/*

		graph = {
		 	'A' : ['B', 'C', 'D'],
		 	'B' : ['A', 'C', 'D'],
		 	'C' : ['A', 'B', 'E'],
		 	'D' : ['A', 'B', 'E', 'F'],
		 	'E' : ['C', 'D', 'F'],
		 	'F' : ['D', 'E'],
		}

	*/

	g := newGraph()
	s := &Stack{}

	g.Insert("A", "B")
	g.Insert("A", "C")
	g.Insert("A", "D")
	g.Insert("B", "C")
	g.Insert("B", "D")
	g.Insert("D", "E")
	g.Insert("D", "F")
	g.Insert("E", "C")
	g.Insert("E", "F")

	s.DFS_Spanning_Tree(g.graph, "A")
	fmt.Println(s)
}
