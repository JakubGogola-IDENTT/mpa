package graph

// Vertex defines coordinates of single graph's vertex on R^2.
type Vertex struct {
	X float64
	Y float64
}

// Edge defines edge of graph's. It contains reference to vertices v,w connected by egde.
type Edge struct {
	V Vertex
	W Vertex
}

// Graph defines graph structure and contains set of verices and edges of this graph.
type Graph struct {
	Vertices []Vertex
	Edges    []Edge
}
