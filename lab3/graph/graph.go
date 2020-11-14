package graph

// // New creates new Graph structure.
// func (g *Graph) New() {
// 	g.Vertices = make([]Vertex)
// }

// AddVertex creates a new vertex with given coords and adds it to graph.
func (g *Graph) AddVertex(x, y float64) (v Vertex) {
	v = Vertex{x, y}
	g.Vertices = append(g.Vertices, v)
	return v
}

// AddEdge create a new edge between given vertices and adds it to graph.
func (g *Graph) AddEdge(v, w Vertex) (e Edge) {
	e = Edge{v, w}
	g.Edges = append(g.Edges, e)
	return e
}
