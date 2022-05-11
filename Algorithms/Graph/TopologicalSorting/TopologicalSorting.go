package main

import "fmt"

type Graph struct {
	g      map[int][]int
	Vertex int
}

func (this *Graph) addEdge(u, v int) {
	this.g[u] = append(this.g[u], v)
}

func (this *Graph) DFS(v int, visited []bool, stack *[]int) {
	visited[v] = true

	for _, i := range this.g[v] {
		if visited[i] == false {
			this.DFS(i, visited, stack)
		}
	}

	*stack = append(*stack, v)
}

func (this *Graph) TopologicalSort() {
	visited := []bool{}
	stack := []int{}
	for i := 0; i < this.Vertex; i++ {
		visited = append(visited, false)
	}

	for i := 0; i < this.Vertex; i++ {
		if visited[i] == false {
			this.DFS(i, visited, &stack)
		}
	}

	fmt.Println(stack)
}

func main() {
	g := &Graph{g: map[int][]int{}}
	g.Vertex = 6
	g.addEdge(5, 2)
	g.addEdge(5, 0)
	g.addEdge(4, 0)
	g.addEdge(4, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 1)

	g.TopologicalSort()
}
