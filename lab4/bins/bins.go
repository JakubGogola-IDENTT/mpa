package bins

// SetBallsAndBins sets number of balls to m and number of bins to n
func (b *Bins) SetBallsAndBins(m, n int) {
	b.m = m
	b.n = n

	b.bins = make(map[int][]int, n)
}

// GetBins returns actual assignment of balls to bins
func (b *Bins) GetBins() map[int][]int {
	return b.bins
}

// Assign randomly puts balls into bins
func (b *Bins) Assign() {
	for i := 0; i < b.m; i++ {
		next := nextInt(b.n)

		b.bins[next] = append(b.bins[next], i)
	}
}
