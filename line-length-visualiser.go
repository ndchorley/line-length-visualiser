package main

import (
	"fmt"
	"os"
	"strings"
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

func makeReport(stats Statistics) string {
	report := fmt.Sprintf("Mean: %f", stats.mean)
	report += "\n"
	report += fmt.Sprintf("Standard deviation: %f", stats.standardDeviation)

	return report
}

func makeVisualisation(lengths []int) string {
	var visualisation strings.Builder

	for _, length := range lengths {
		for i := 1; i <= length; i++ {
			visualisation.WriteString(`◼`)
		}

		visualisation.WriteString("\n")
	}

	return visualisation.String()
}

func toLengths(lines []string) []int {
	lengths := make([]int, len(lines))

	for index, line := range lines {
		lengths[index] = len(line)
	}

	return lengths
}
