package analytics

import (
	"bufio"
	"flag"
	"fmt"
	alg "lab1/algorithms"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (a *Analytics) parseFlags() {
	flag.StringVar(&a.algorithmType, "t", alg.QUICK, "type of sorting algorithm")

	flag.IntVar(&a.lowerBound, "l", 100, "lower bound of input size")
	flag.IntVar(&a.upperBound, "u", 10000, "upper bound of input size")
	flag.IntVar(&a.step, "s", 100, "step size")
	flag.IntVar(&a.repetitions, "r", 1000, "number of repetitions")

	flag.Parse()
}

func (a *Analytics) newFileWithWriter() (f *os.File, w *bufio.Writer) {
	fileName := fmt.Sprintf("%s_%d_%d_%d_%d.csv", a.algorithmType, a.lowerBound, a.upperBound, a.step, a.repetitions)
	f, err := os.Create(fileName)

	checkError(err)

	w = bufio.NewWriter(f)

	return f, w
}

// RunTests runs algorithms tests
func (a *Analytics) RunTests() {
	a.parseFlags()

	f, w := a.newFileWithWriter()
	defer f.Close()

	_, err := w.WriteString("input_size,comparisons_num\n")
	checkError(err)

	alg := alg.Algorithm{}

	for i := a.lowerBound; i <= a.upperBound; i += a.step {
		alg.SetSize(i)

		for j := 0; j < a.repetitions; j++ {
			alg.Run(a.algorithmType)

			row := fmt.Sprintf("%d,%d\n", i, alg.Comps)

			_, err := w.WriteString(row)
			checkError(err)
		}
	}

	w.Flush()
}
