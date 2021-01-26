package analytics

import (
	"bufio"
	"log"
	"os"
)

func createFileWithWriter(fileName string) (*os.File, *bufio.Writer) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	return f, w
}
