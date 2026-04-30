package main

import (
	"fmt"
	"os"
)

func main() {
	fileNames := os.Args[1:]
	lines := readLines(fileNames)

	lengths := toLengths(lines)
	stats := gatherStatistics(lengths)

	fmt.Println(visualisationFor(lengths))
	fmt.Println(reportFor(stats))
}
