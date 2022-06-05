package main

import "fmt"

type Graph struct {
	g      [][]int
	Vertex int
}

func (this *Graph) addEdge(v, w int) {
	this.g[v] = append(this.g[v], w)
	this.g[w] = append(this.g[w], v)
}

func (this *Graph) GraphColoring() {
	result := []int{}
	available := []bool{}
	for i := 0; i < this.Vertex; i++ {
		result = append(result, -1)
		available = append(available, false)
	}

	result[0] = 0

	for u := 0; u < this.Vertex; u++ {
		for _, i := range this.g[u] {
			if result[i] != -1 {
				available[result[i]] = true
			}
		}

		color := 0
		for color < this.Vertex {
			if available[color] == false {
				break
			}
			color++
		}

		result[u] = color

		for _, i := range this.g[u] {
			if result[i] != -1 {
				available[result[i]] = false
			}
		}
	}

	for i := 0; i < this.Vertex; i++ {
		fmt.Println("Vertex", i, " ---> Color", result[i])
	}
}

func main() {
	g := Graph{make([][]int, 5), 5}
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.GraphColoring()
}

// https://www.geeksforgeeks.org/graph-coloring-set-2-greedy-algorithm/
