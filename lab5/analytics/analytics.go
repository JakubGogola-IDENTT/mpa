package analytics

import (
	"bufio"
	"flag"
	"fmt"
	"lab5/trees"
	"log"
	"os"
)

// Init sets testing params from flags
func (a *Analytics) Init() {
	flag.IntVar(&a.lowerBound, "l", 100, "lower size of input")
	flag.IntVar(&a.upperBound, "u", 10000, "upper size of input")
	flag.IntVar(&a.step, "s", 100, "step of increasing input size")
	flag.IntVar(&a.repetitions, "r", 1000, "number of repetitions for given input size")

	flag.Parse()
}

// RunTests performs algorithm tests with given params
func (a *Analytics) RunTests() {
	f, err := os.Create("results.csv")

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	defer f.Close()

	_, err = w.WriteString("n,min,max\n")

	if err != nil {
		log.Fatal(err)
	}

	t := trees.Tree{}

	for n := a.lowerBound; n <= a.upperBound; n += a.step {
		if n%100 == 0 {
			fmt.Printf("Progress: %d\n", n)
		}

		for r := 0; r < a.repetitions; r++ {
			t.Generate(n)

			min, max := t.Stats()

			_, err := w.WriteString(fmt.Sprintf("%d,%d,%d\n", n, min, max))

			if err != nil {
				log.Fatal(err)
			}
		}

		w.Flush()
	}
}
