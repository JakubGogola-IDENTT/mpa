package delaunay

import (
	"math/rand"
	"time"
)

// New initializes all parameters needed to run Delaunay triangulation algorithm.
func (d *Delaunay) New() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// SetInputSize sets number of random points to triangulate.
func (d *Delaunay) SetInputSize(n int) {
	d.n = n
}
