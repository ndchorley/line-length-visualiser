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

	visualisation := makeVisualisation(lengths)
	report := makeReport(stats)

	fmt.Println(visualisation)
	fmt.Println(report)
}

func toLengths(lines []string) []int {
	lengths := make([]int, len(lines))

	for index, line := range lines {
		lengths[index] = len(line)
	}

	return lengths
}
