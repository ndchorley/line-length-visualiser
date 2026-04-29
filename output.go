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
	binLabels := binLabelsFor(histogram)

	labelled := make([]string, 0)

	for i, line := range lines {
		thisLine := binLabels[i] + "|" + line

		labelled = append(labelled, thisLine)
	}

	return labelled
}

func binLabelsFor(histogram []Bin) []string {
	labels := make([]string, 0)

	for _, bin := range histogram {
		label := fmt.Sprintf("%d-%d", bin.start, bin.end)

		labels = append(labels, label)
	}

	return addPadding(labels)
}

func addPadding(labels []string) []string {
	max := maxLengthOf(labels)

	withPadding := make([]string, 0)

	for _, label := range labels {
		if len(label) < max {
			padded := padTo(max, label)

			withPadding = append(withPadding, padded)
		} else {
			withPadding = append(withPadding, label)
		}
	}

	return withPadding
}

func padTo(max int, label string) string {
	amountOfPadding := max - len(label)

	padding := strings.Repeat(" ", amountOfPadding)

	return padding + label
}

func maxLengthOf(labels []string) int {
	maxSoFar := len(labels[0])

	for i := 1; i < len(labels); i++ {
		thisLength := len(labels[i])

		if thisLength > maxSoFar {
			maxSoFar = thisLength
		}
	}

	return maxSoFar
}

func countsFrom(histogram []Bin) []int {
	counts := make([]int, 0)

	for _, bin := range histogram {
		counts = append(counts, bin.count)
	}

	return counts
}
