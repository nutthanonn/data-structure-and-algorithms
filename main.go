/*

---- Kruskal Algorithm ----


*/

package main

import "fmt"

type Graph struct {
	graph map[string]map[string]int
}

func main() {
	g := &Graph{graph: make(map[string]map[string]int)}
	fmt.Println(g)
}
