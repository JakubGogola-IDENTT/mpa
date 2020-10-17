package algorithms

// Algorithm structure contains current algorithm params and stats
type Algorithm struct {
	Comps     int
	ArraySize int
	Array     []int
}

const (
	// MERGE = merge sort algorithm
	MERGE string = "merge"
	// QUICK = quick sort algorithm
	QUICK string = "quick"
)
