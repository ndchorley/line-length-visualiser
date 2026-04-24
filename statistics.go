package main

import (
	"math"
	"slices"
)

type Statistics struct {
	numberOfLines     int
	mean              float64
	standardDeviation float64
	histogram         map[Bin]int
}

type Bin struct {
	start int
	end   int
}

func gatherStatistics(lengths []int) Statistics {
	mean := calculateMean(lengths)
	standardDeviation :=
		calculateStandardDeviation(mean, lengths)

	histogram := buildHistogram(lengths)

	return Statistics{len(lengths), mean, standardDeviation, histogram}
}

func calculateMean(lengths []int) float64 {
	total := 0.0

	for _, length := range lengths {
		total += float64(length)
	}

	return total / float64(len(lengths))
}

func calculateStandardDeviation(mean float64, lengths []int) float64 {
	totalSquaredDeviation := 0.0

	for _, length := range lengths {
		totalSquaredDeviation += (float64(length) - mean) * (float64(length) - mean)
	}

	return math.Sqrt(totalSquaredDeviation / float64(len(lengths)))
}

func buildHistogram(lengths []int) map[Bin]int {
	lengthsRange := slices.Max(lengths) - slices.Min(lengths)
	numberOfBins := 10
	binWidth := int(math.Ceil(float64(lengthsRange) / float64(numberOfBins)))

	histogram := make(map[Bin]int)

	for i := 1; i <= numberOfBins; i++ {
		start := slices.Min(lengths) + binWidth*(i-1)
		end := slices.Min(lengths) + binWidth*i

		bin := Bin{start, end}
		histogram[bin] = countHowMany(bin, lengths)
	}

	return histogram
}

func countHowMany(bin Bin, lengths []int) int {
	count := 0

	for _, length := range lengths {
		if length >= bin.start && length < bin.end {
			count += 1
		}
	}

	return count
}
