package analytics

import (
	"fmt"
	perm "lab2/permutations"
)

// RunTests runs tests
func (a *Analytics) RunTests() {
	a.parseFlags()

	p := perm.Permutation{}

	f, w := a.newFileWithWriter()
	defer f.Close()
	defer w.Flush()

	_, err := w.WriteString("n,cycles,records,fixed_points\n")
	checkError(err)

	for n := a.start; n <= a.end; n += a.step {
		p.SetSize(n)

		for r := 0; r < a.repeats; r++ {
			p.Permute()
			cycles, records, fixedPoints := p.Properties()

			row := fmt.Sprintf("%d,%d,%d,%d\n", n, cycles, records, fixedPoints)

			_, err := w.WriteString(row)
			checkError(err)
		}
	}
}
