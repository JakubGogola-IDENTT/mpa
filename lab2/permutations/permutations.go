package permutations

import "fmt"

// Init initializes permutation
func (p *Permutation) Init() {

}

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

	fmt.Println(unpermuted)

	for i := range perm {
		idx := nextInt(len(unpermuted))

		perm[i] = unpermuted[idx]

		unpermuted = remove(unpermuted, idx)
	}

	p.Perm = perm
}

// Cycles returns number of cycles in permutation
func (p *Permutation) Cycles() {

}
