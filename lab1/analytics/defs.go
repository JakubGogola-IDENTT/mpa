package analytics

// Analytics struct contains info about currently tested algorithm
type Analytics struct {
	algorithmType string
	lowerBound    int
	upperBound    int
	step          int
	repetitions   int
}
