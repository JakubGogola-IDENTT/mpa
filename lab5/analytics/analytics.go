package analytics

import (
	"flag"
	"fmt"
	"lab5/trees"
	"log"
)

// Init sets testing params from flags
func (a *Analytics) Init() {
	flag.IntVar(&a.lowerBound, "l", 100, "lower size of input")
	flag.IntVar(&a.upperBound, "u", 10000, "upper size of input")
	flag.IntVar(&a.step, "s", 100, "step of increasing input size")
	flag.IntVar(&a.repetitions, "r", 1000, "number of repetitions for given input size")

	flag.Parse()
}

// TestForDifferentN performs algorithm tests with given params
func (a *Analytics) TestForDifferentN() {
	fmt.Println("Test: different n's")

	f, w := createFileWithWriter("results.csv")
	defer f.Close()

	t := &trees.Tree{}

	_, err := w.WriteString("n,min,max\n")

	if err != nil {
		log.Fatal(err)
	}

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

// TestUniformDistribution test if trees for given n are generated with uniform distribution
func (a *Analytics) TestUniformDistribution() {
	fmt.Println("Test: uniform distribution")

	f, w := createFileWithWriter("results_uniform.csv")
	defer f.Close()

	t := &trees.Tree{}

	counter := make(map[string]int)

	for r := 0; r < 10000000; r++ {
		if r%100000 == 0 {
			fmt.Printf("Progress: %d\n", r)
		}

		t.Generate(4)
		counter[t.SeqToString()]++
	}

	_, err := w.WriteString("k,v\n")

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range counter {
		_, err := w.WriteString(fmt.Sprintf("%s,%d\n", k, v))

		if err != nil {
			log.Fatal(err)
		}

		w.Flush()
	}
}

// TestProtecionNumber test protection number for given n
func (a *Analytics) TestProtecionNumber() {
	fmt.Println("Test: protection number")

	f, w := createFileWithWriter("results_protection.csv")
	defer f.Close()

	t := &trees.Tree{}

	_, err := w.WriteString("k,v\n")

	if err != nil {
		log.Fatal(err)
	}

	counter := make(map[int]int)

	for r := 0; r < 1000000; r++ {
		if r%100000 == 0 {
			fmt.Printf("Progress: %d\n", r)
		}

		t.Generate(500)
		min, _ := t.Stats()

		counter[min]++
	}

	for k, v := range counter {
		_, err := w.WriteString(fmt.Sprintf("%d,%d\n", k, v))

		if err != nil {
			log.Fatal(err)
		}

		w.Flush()
	}
}
