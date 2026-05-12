package main

import (
	"os"
	"slices"
	"strings"
)

func readLinesIn(files []string) []string {
	allLines := make([]string, 0)

	for _, file := range files {
		theseLines := linesIn(file)

		allLines = slices.Concat(allLines, theseLines)
	}

	return allLines
}

func linesIn(file string) []string {
	contents := readFile(file)

	lines := strings.Split(contents, "\n")

	return lines[0 : len(lines)-1]
}

func readFile(file string) string {
	bytes, _ := os.ReadFile(file)
	contents := string(bytes)

	return contents
}
