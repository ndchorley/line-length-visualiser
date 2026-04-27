package main

import (
	"fmt"
	"strings"
)

func reportFor(stats Statistics) string {
	report := fmt.Sprintf("Number of lines: %d", stats.numberOfLines)
	report += "\n\n"
	report += fmt.Sprintf("Mean: %f", stats.mean)
	report += "\n"
	report += fmt.Sprintf("Standard deviation: %f", stats.standardDeviation)
	report += "\n\n"
	report += fmt.Sprint(stats.histogram)

	return report
}

func visualisationFor(lengths []int) string {
	lines := linesFor(lengths)
	visualisation := strings.Join(lines, "\n")

	return visualisation
}

func linesFor(lengths []int) []string {
	lines := make([]string, 0)

	for _, length := range lengths {
		var line strings.Builder

		for i := 1; i <= length; i++ {
			line.WriteString(`◼`)
		}

		lines = append(lines, line.String())
	}
	return lines
}
