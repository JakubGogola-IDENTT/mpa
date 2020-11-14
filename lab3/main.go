package main

import (
	"fmt"
	"lab3/graph"
)

func main() {
	g := graph.Graph{}

	v := g.AddVertex(2.3, 4.5)
	w := g.AddVertex(1.2, 3.3)
	fmt.Println(g.Vertices)

	g.AddEdge(v, w)
	fmt.Println(g.Edges)
}
