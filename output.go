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
	report += graphOf(stats.histogram)

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

func graphOf(histogram []Bin) string {
	counts := countsFrom(histogram)
	graphLines := addBinLabels(linesFor(counts), histogram)

	return fmt.Sprint(strings.Join(graphLines, "\n"))
}

func addBinLabels(lines []string, histogram []Bin) []string {
	linesWithLabels := make([]string, 0)

	for index, line := range lines {
		lineWithLabel := fmt.Sprintf("%d-%d  |%s", histogram[index].start, histogram[index].end, line)

		linesWithLabels = append(linesWithLabels, lineWithLabel)
	}

	return linesWithLabels
}

func countsFrom(histogram []Bin) []int {
	counts := make([]int, 0)

	for _, bin := range histogram {
		counts = append(counts, bin.count)
	}

	return counts
}
