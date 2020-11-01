package analytics

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (a *Analytics) parseFlags() {
	flag.IntVar(&a.start, "start", 100, "lower bound of input size")
	flag.IntVar(&a.end, "end", 10000, "upper bound of input size")
	flag.IntVar(&a.step, "step", 100, "input size step")
	flag.IntVar(&a.repeats, "repeats", 500, "number of repeats for each input size")

	flag.Parse()
}

func (a *Analytics) newFileWithWriter() (f *os.File, w *bufio.Writer) {
	fileName := fmt.Sprintf("perms_%d_%d_%d_%d.csv", a.start, a.end, a.step, a.repeats)
	f, err := os.Create(fileName)

	checkError(err)

	w = bufio.NewWriter(f)

	return f, w
}
