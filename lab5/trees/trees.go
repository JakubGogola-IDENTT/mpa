package trees

// Generate creates random tree of given number of nodes
func (t *Tree) Generate(n int) {
	seq := randomArray(n)

	var prefix, sufix, word []int
	var partialSum = 0

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
