/*

Bellman Ford Algorithm Dynamic Programming

    graph = {
        '1': {'2': 6, '3': 5, '4': 5},
        '2': {'5': -1},
        '3': {'2': -2, '5': 1},
        '4': {'3': -2, '6': -1},
        '5': {'7': 3},
        '6': {'7': 3},
        '7': {},
    }

*/

package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertex int
	g      [][]int
}

func singleSourceShortestPath(arr []int) {
	fmt.Println("Vertex\t\tShortest Path")
	for i := 0; i < len(arr); i++ {
		if arr[i] != int(math.Inf(1)) {
			fmt.Printf("%d\t\t%d\n", i, arr[i])
		} else {
			fmt.Printf("%d\t\t%s\n", i, "∞")
		}
	}
}

func (this *Graph) addEdge(u, v, w int) {
	this.g = append(this.g, []int{u, v, w})
}

func (this *Graph) BellmanFord(start int) {
	dist := []int{}
	for i := 0; i < this.Vertex; i++ {
		dist = append(dist, int(math.Inf(1)))
	}

	dist[start] = 0

	for r := 0; r < this.Vertex-1; r++ {
		for i := 0; i < this.Vertex; i++ {
			u := this.g[i][0]
			v := this.g[i][1]
			w := this.g[i][2]

			if dist[u] != int(math.Inf(1)) && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
			}
		}
	}

	for i := 0; i < this.Vertex; i++ {
		u := this.g[i][0]
		v := this.g[i][1]
		w := this.g[i][2]
		if dist[u] != int(math.Inf(1)) && dist[u]+w < dist[v] {
			fmt.Println("Graph contains negative weight cycle")
			return
		}
	}

	singleSourceShortestPath(dist)
}

func main() {
	g := &Graph{Vertex: 7}
	g.addEdge(1, 2, 6)
	g.addEdge(1, 3, 5)
	g.addEdge(1, 4, 5)
	g.addEdge(2, 5, -1)
	g.addEdge(3, 2, -2)
	g.addEdge(3, 5, 1)
	g.addEdge(4, 3, -2)
	g.addEdge(4, 6, -1)
	g.addEdge(5, 7, 3)
	g.addEdge(6, 7, 3)

	g.BellmanFord(1)
}
