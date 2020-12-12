package bins

// Bins defines set of bins and balls
type Bins struct {
	// m is balls count
	m int
	// n is bins count
	n    int
	bins map[int][]int
}
