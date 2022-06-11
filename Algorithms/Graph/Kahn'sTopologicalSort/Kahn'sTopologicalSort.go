package main

type Topo struct {
	Graph  map[int][]int
	Vertex int
}

func (this *Topo) addEdge(v, w int) {
	this.Graph[v] = append(this.Graph[v], w)
}

func (this *Topo) TopologicalSort() []int {
	indeg := []int{}
	for i := 0; i < this.Vertex; i++ {
		indeg = append(indeg, 0)
	}

	for _, i := range this.Graph {
		for _, j := range i {
			indeg[j]++
		}
	}

	q := []int{}
	stack := []int{}
	checkCycle := 0
	for i := 0; i < this.Vertex; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		u := q[0]
		stack = append(stack, u)
		q = q[1:]

		for _, v := range this.Graph[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if checkCycle != this.Vertex {
		return []int{}
	} else {
		return stack
	}

}

func main() {

}
