package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	lines := readLines(fileName)

	lengths := toLengths(lines)
	stats := gatherStatistics(lengths)

	fmt.Println(visualisationFor(lengths))
	fmt.Println(reportFor(stats))
}
