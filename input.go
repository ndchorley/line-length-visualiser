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

func linesIn(fileName string) []string {
	contents := readFile(fileName)

	lines := strings.Split(contents, "\n")

	return lines[0 : len(lines)-1]
}

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	contents := string(bytes)

	return contents
}
