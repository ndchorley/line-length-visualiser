package main

func toLengths(lines []string) []int {
	lengths := make([]int, len(lines))

	for index, line := range lines {
		if len(line) != 0 {
			lengths[index] = len(line)
		}
	}

	return lengths
}
