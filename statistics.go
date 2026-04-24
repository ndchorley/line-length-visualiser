package main

import (
	"math"
)

type Statistics struct {
	numberOfLines     int
	mean              float64
	standardDeviation float64
}

func gatherStatistics(lengths []int) Statistics {
	mean := calculateMean(lengths)
	standardDeviation :=
		calculateStandardDeviation(mean, lengths)

	return Statistics{len(lengths), mean, standardDeviation}
}

func calculateStandardDeviation(mean float64, lengths []int) float64 {
	totalDeviation := 0.0

	for _, length := range lengths {
		totalDeviation += (float64(length) - mean) * (float64(length) - mean)
	}

	return math.Sqrt(totalDeviation / float64(len(lengths)))
}

func calculateMean(lengths []int) float64 {
	total := 0.0

	for _, length := range lengths {
		total += float64(length)
	}

	return total / float64(len(lengths))
}
