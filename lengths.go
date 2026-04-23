package main

func toLengths(lines []string) []int {
	lengths := make([]int, len(lines))

	for index, line := range lines {
		lengths[index] = len(line)
	}

	return lengths
}
