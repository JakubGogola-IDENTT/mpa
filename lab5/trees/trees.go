package trees

// Generate creates random tree of given number of nodes
func (t *Tree) Generate(n int) {
	seq := randomArray(n)

	var prefix, sufix, word []int
	partialSum := 0

	for _, s := range seq {
		word = append(word, s)

		partialSum += s

		if partialSum != 0 {
			continue
		}

		if s == -1 {
			prefix = append(prefix, word...)
		} else {
			prefix = append(prefix, 1)

			var nextSufix []int
			nextSufix = append(nextSufix, -1)

			for _, w := range word[1 : len(word)-1] {
				nextSufix = append(nextSufix, -w)
			}

			nextSufix = append(nextSufix, sufix...)
			sufix = nextSufix
		}

		word = nil
	}

	t.sequence = append(prefix, sufix...)
}

// IsValid checks if generated sequence of braces represents valid binary tree.
func (t *Tree) IsValid() bool {
	partialSum := 0

	for _, s := range t.sequence {
		if s != -1 && s != 1 {
			return false
		}

		partialSum += s

		if partialSum < 0 {
			return false
		}
	}

	return partialSum == 0
}
