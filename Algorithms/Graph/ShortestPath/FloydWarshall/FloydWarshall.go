package main

/*

* NOTE *

node = key
edge = value

graph = {
    "1": {"2": 3, "4": 7},
    "2": {"1": 8, "3": 2},
    "3": {"1": 5, "4": 1},
    "4": {"1": 2},
}



*/

import (
	"fmt"
	"math"
	"strconv"
)

type Graph struct {
	graph map[string]map[string]int
}

func makeAdjacencyMatrix(g *Graph) [][]int {
	final := [][]int{}
	l := len(g.graph) + 1
	for i := 1; i < l; i++ {
		val := []int{}
		for j := 1; j < l; j++ {
			if i == j {
				val = append(val, 0)
				continue
			}

			if g.graph[strconv.Itoa(i)][strconv.Itoa(j)] == 0 {
				val = append(val, int(math.Inf(1)))
				continue
			}

			val = append(val, g.graph[strconv.Itoa(i)][strconv.Itoa(j)])
		}

		final = append(final, val)
	}

	return final
}

func FloydWarshall(adjacencyMatrix [][]int) [][]int {
	n := len(adjacencyMatrix)

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if adjacencyMatrix[i][k] == int(math.Inf(1)) || adjacencyMatrix[k][j] == int(math.Inf(1)) {
					continue
				}

				val := adjacencyMatrix[i][k] + adjacencyMatrix[k][j]
				current := adjacencyMatrix[i][j]

				if current > val {
					adjacencyMatrix[i][j] = val
				}
			}
		}
	}
	return adjacencyMatrix
}

func main() {
	g := &Graph{graph: make(map[string]map[string]int)}
	g.graph["1"] = map[string]int{"2": 3, "4": 7}
	g.graph["2"] = map[string]int{"1": 8, "3": 2}
	g.graph["3"] = map[string]int{"1": 5, "4": 1}
	g.graph["4"] = map[string]int{"1": 2}

	adjacencyMatrix := makeAdjacencyMatrix(g)
	fmt.Println(adjacencyMatrix)
	fmt.Println("-----------------------------------")
	AllPairsShortestPath := FloydWarshall(adjacencyMatrix)
	fmt.Println(AllPairsShortestPath)
}
