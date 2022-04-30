package main

import (
	"fmt"
	"sort"
)

type Graph struct {
	vertex int
	graph  [][]int
	MST    [][]int
}

func (this *Graph) add_edge(u, v, w int) {
	this.graph = append(this.graph, []int{u, v, w})
}

func (this *Graph) find(parrent []int, i int) int {
	if parrent[i] == i {
		return i
	}
	return this.find(parrent, parrent[i])
}

func (this *Graph) unioun(parent []int, rank []int, x, y int) {
	xroot := this.find(parent, x)
	yroot := this.find(parent, y)

	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}

}

func (this *Graph) kruskal() {
	res := [][]int{}
	i, e := 0, 0
	sort.Slice(this.graph, func(i, j int) bool { return this.graph[i][2] < this.graph[j][2] })
	parent := []int{}
	rank := []int{}

	for i := 0; i < this.vertex; i++ {
		parent = append(parent, i)
		rank = append(rank, 0)
	}

	for e < this.vertex-1 {
		u := this.graph[i][0]
		v := this.graph[i][1]
		w := this.graph[i][2]

		i++
		x := this.find(parent, u)
		y := this.find(parent, v)
		if x != y {
			e++
			res = append(res, []int{u, v, w})
			this.unioun(parent, rank, x, y)
		}
	}
	this.MST = res
}

func main() {
	g := &Graph{vertex: 6}
	g.add_edge(0, 1, 4)
	g.add_edge(0, 2, 4)
	g.add_edge(1, 2, 2)
	g.add_edge(1, 0, 4)
	g.add_edge(2, 0, 4)
	g.add_edge(2, 1, 2)
	g.add_edge(2, 3, 3)
	g.add_edge(2, 5, 2)
	g.add_edge(2, 4, 4)
	g.add_edge(3, 2, 3)
	g.add_edge(3, 4, 3)
	g.add_edge(4, 2, 4)
	g.add_edge(4, 3, 3)
	g.add_edge(5, 2, 2)
	g.add_edge(5, 4, 3)
	g.kruskal()

	for i := 0; i < g.vertex-1; i++ {
		fmt.Printf("Edge %d to Edge %d -> weight: %d\n", g.MST[i][0], g.MST[i][1], g.MST[i][2])
	}

}
