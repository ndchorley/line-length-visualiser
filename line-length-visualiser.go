package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	lines := readLinesIn(files)

	lengths := lengthsOf(lines)
	stats := gatherStatistics(lengths)

	if len(files) == 1 {
		fmt.Println(visualisationFor(lengths))
	}

	fmt.Println(reportFor(stats))
}
