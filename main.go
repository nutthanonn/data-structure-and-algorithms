// package main

// import "fmt"

// type Graph struct {
// 	g map[string][]string
// }

// type Stack struct {
// 	s []string
// }

// func createGraph() *Graph {
// 	return &Graph{g: make(map[string][]string)}
// }

// func (this *Graph) addEdge(u, v string) {
// 	this.g[u] = append(this.g[u], v)
// 	this.g[v] = append(this.g[v], u)
// }

// func (this *Stack) DFS(g *Graph, start string) {
// 	check := true
// 	for _, v := range this.s {
// 		if v == start {
// 			check = false
// 		}
// 	}

// 	if check {
// 		this.s = append(this.s, start)
// 		for _, n := range g.g[start] {
// 			this.DFS(g, n)
// 		}
// 	}
// }

// func main() {
// 	g := createGraph()
// 	s := Stack{}
// 	g.addEdge("A", "B")
// 	g.addEdge("A", "S")
// 	g.addEdge("S", "C")
// 	g.addEdge("S", "G")
// 	g.addEdge("C", "D")
// 	g.addEdge("C", "E")
// 	g.addEdge("C", "F")
// 	g.addEdge("E", "H")
// 	g.addEdge("H", "G")
// 	fmt.Println(g)
// 	s.DFS(g, "A")

// 	fmt.Println(s.s)
// }

package main

import (
	"fmt"
)

type Graph struct {
	g map[string][]string
}

type Queue struct {
	Q    []string
	curr int
	node int
}

func createGraph() *Graph {
	return &Graph{g: make(map[string][]string)}
}

func (this *Graph) addEdge(u, v string) {
	this.g[u] = append(this.g[u], v)
	this.g[v] = append(this.g[v], u)
}

func isInArr(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}

	return false
}

func (this *Queue) BFS(g *Graph, start string) {
	if len(this.Q) == 0 {
		this.Q = append(this.Q, start)
	}

	if len(this.Q) == this.node {
		return
	}

	for _, v := range g.g[start] {
		if isInArr(this.Q, v) {
			continue
		}
		this.Q = append(this.Q, v)
	}

	this.curr++
	this.BFS(g, this.Q[this.curr])
}

func main() {
	g := createGraph()
	g.addEdge("0", "9")
	g.addEdge("0", "7")
	g.addEdge("0", "11")
	g.addEdge("9", "10")
	g.addEdge("9", "8")
	g.addEdge("10", "1")
	g.addEdge("8", "1")
	g.addEdge("8", "12")
	g.addEdge("7", "3")
	g.addEdge("7", "6")
	g.addEdge("6", "5")
	g.addEdge("3", "2")
	g.addEdge("3", "4")
	g.addEdge("12", "2")

	q := Queue{node: 13}
	fmt.Println(g)
	q.BFS(g, "0")
	fmt.Println(q.Q)
	fmt.Println(q.curr)
}
