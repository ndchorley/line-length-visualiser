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

type Statistics struct {
	mean              float64
	standardDeviation float64
}

func makeReport(stats Statistics) string {
	report := fmt.Sprintf("Mean: %f", stats.mean)
	report += "\n"
	report += fmt.Sprintf("Standard deviation: %f", stats.standardDeviation)

	return report
}

func gatherStatistics(lengths []int) Statistics {
	mean := calculateMean(lengths)
	standardDeviation :=
		calculateStandardDeviation(mean, lengths)

	return Statistics{mean, standardDeviation}
}

func calculateStandardDeviation(mean float64, lengths []int) float64 {
	totalDeviation := 0.0

	for _, length := range lengths {
		totalDeviation += (float64(length) - mean) * (float64(length) - mean)
	}

	return totalDeviation / float64(len(lengths))
}

func calculateMean(lengths []int) float64 {
	total := 0.0

	for _, length := range lengths {
		total += float64(length)
	}

	return total / float64(len(lengths))
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

func readLines(fileName string) []string {
	contents := readFile(fileName)

	lines := strings.Split(contents, "\n")

	return lines[0 : len(lines)-1]
}

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	contents := string(bytes)

	return contents
}

func toLengths(lines []string) []int {
	lengths := make([]int, len(lines))

	for index, line := range lines {
		lengths[index] = len(line)
	}

	return lengths
}
