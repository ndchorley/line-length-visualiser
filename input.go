package main

import (
	"os"
	"slices"
	"strings"
)

func readLinesIn(fileNames []string) []string {
	lines := make([]string, 0)

	for _, fileName := range fileNames {
		theseLines := linesIn(fileName)

		lines = slices.Concat(lines, theseLines)
	}

	return lines
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
