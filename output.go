package main

import (
	"fmt"
	"strings"
)

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
