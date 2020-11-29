package graph

// // New creates new Graph structure.
// func (g *Graph) New() {
// 	g.Vertices = make([]Vertex)
// }

// AddVertex creates a new vertex with given coords and adds it to graph.
func (g *Graph) AddVertex(x, y float64) (v Vertex) {
	v = Vertex{X: x, Y: y}
	g.Vertices = append(g.Vertices, v)
	return v
}

// AddEdge create a new edge between given vertices and adds it to graph.
func (g *Graph) AddEdge(v, w Vertex) (e Edge) {
	e = Edge{V: v, W: w}
	g.Edges = append(g.Edges, e)
	return e
}

// DeleteEdge deletes edge between given vertices
func (g *Graph) DeleteEdge(v, w Vertex) {

}
