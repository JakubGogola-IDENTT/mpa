package graph

// IsEqual checks if coords of current vertex v are equal to coords of given w.
func (v *Vertex) IsEqual(w Vertex) bool {
	return v.X == w.X && v.Y == w.Y
}

// CompareX compares x coord of current vertex with given x coord.
func (v *Vertex) CompareX(x float64) int {
	if v.X > x {
		return 1
	}

	if v.X < x {
		return -1
	}

	return 0
}

// CompareY compares y coord of current vertex with given y coord.
func (v *Vertex) CompareY(y float64) int {
	if v.Y > y {
		return 1
	}

	if v.Y < y {
		return -1
	}

	return 0
}
