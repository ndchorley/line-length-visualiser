package main

import (
	"fmt"
	"os"
)

func main() {
	fileNames := os.Args[1:]
	lines := readLines(fileNames)

	lengths := lengthsOf(lines)
	stats := gatherStatistics(lengths)

	if len(fileNames) == 1 {
		fmt.Println(visualisationFor(lengths))
	}

	fmt.Println(reportFor(stats))
}
