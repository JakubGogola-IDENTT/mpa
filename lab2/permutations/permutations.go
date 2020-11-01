package permutations

// Init initializes permutation
func (p *Permutation) Init() {

}

// SetSize sets size of permutation
func (p *Permutation) SetSize(size int) {
	p.Size = size
}

// Permute permutes permutation
func (p *Permutation) Permute() {
	p.Perm = make([][]int, p.Size)

	for i := range p.Perm {
		p.Perm[i] = make([]int, 2)

		value := i + 1

		p.Perm[i][0], p.Perm[i][1] = value, value
	}

	for i := range p.Perm {
		j := nextInt(p.Size)

		p.Perm[i][1], p.Perm[j][1] = p.Perm[j][1], p.Perm[i][1]
	}
}
