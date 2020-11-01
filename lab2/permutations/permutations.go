package permutations

// SetSize sets size of permutation
func (p *Permutation) SetSize(size int) {
	p.Size = size
}

// Permute permutes permutation
func (p *Permutation) Permute() {
	perm := make([]int, p.Size)

	for i := range perm {
		perm[i] = i
	}

	unpermuted := make([]int, p.Size)
	copy(unpermuted, perm)

	for i := range perm {
		idx := nextInt(len(unpermuted))

		perm[i] = unpermuted[idx]

		unpermuted = remove(unpermuted, idx)
	}

	p.Perm = perm
}

// Properties returns numbers of cycles, records and fixed points
func (p *Permutation) Properties() (cycles, records, fixedPoints int) {
	visited := make(map[int]bool)
	perm := p.Perm

	for i, value := range perm {
		slice := perm[:i]

		isRecord := true

		for _, v := range slice {
			if v >= value {
				isRecord = false
				break
			}
		}

		if isRecord {
			records++
		}

		if _, ok := visited[i]; ok {
			continue
		}

		start := i
		next := perm[start]
		visited[start] = true

		if start == next {
			fixedPoints++
		}

		for next != start {
			visited[next] = true

			next = perm[next]
		}

		cycles++
	}

	return cycles, records, fixedPoints
}
